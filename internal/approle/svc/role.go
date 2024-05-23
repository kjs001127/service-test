package svc

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/approle/model"
	authgen "github.com/channel-io/ch-app-store/internal/auth/general"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type ClaimsDTO struct {
	NativeClaims model.Claims `json:"nativeClaims"`
	AppClaims    model.Claims `json:"appClaims"`
}

type AppRoleSvc struct {
	roleCli        authgen.RoleFetcher
	roleRepo       AppRoleRepository
	secretRepo     AppSecretRepository
	typeRule       map[model.RoleType]TypeRule
	appSvc         app.AppQuerySvc
	nativeServices []string
}

func NewAppRoleSvc(
	roleCli authgen.RoleFetcher,
	roleRepo AppRoleRepository,
	typeRule map[model.RoleType]TypeRule,
	appSvc app.AppQuerySvc,
	secretRepo AppSecretRepository,
	nativeServices []string,
) *AppRoleSvc {
	return &AppRoleSvc{
		roleCli:        roleCli,
		roleRepo:       roleRepo,
		typeRule:       typeRule,
		secretRepo:     secretRepo,
		appSvc:         appSvc,
		nativeServices: nativeServices,
	}
}

func (s *AppRoleSvc) CreateRoles(ctx context.Context, appID string) error {
	for _, roleType := range model.AvailableRoleTypes {
		typeRule := s.typeRule[roleType]
		res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
			Claims:                typeRule.DefaultClaimsOf(appID),
			AllowedPrincipalTypes: typeRule.PrincipalTypes,
			AllowedGrantTypes:     typeRule.GrantTypes,
		})
		if err != nil {
			return err
		}

		if err := s.roleRepo.Save(ctx, &model.AppRole{
			AppID:  appID,
			Type:   roleType,
			RoleID: res.Role.Id,
			Credentials: &model.Credentials{
				ClientSecret: res.Credentials.ClientSecret,
				ClientID:     res.Credentials.ClientId,
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppRoleSvc) UpdateRole(ctx context.Context, appID string, roleType model.RoleType, request *ClaimsDTO) error {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, roleType)
	if err != nil {
		return err
	}

	appClaims, err := s.appClaimsToProto(ctx, request.AppClaims)
	if err != nil {
		return err
	}

	nativeClaims, err := s.nativeClaimsToProto(roleType, request.NativeClaims)
	if err != nil {
		return err
	}

	allClaims := make([]*protomodel.Claim, 0, len(request.AppClaims)+len(request.NativeClaims))
	allClaims = append(allClaims, appClaims...)
	allClaims = append(allClaims, nativeClaims...)

	_, err = s.roleCli.UpdateRole(ctx, &service.ReplaceRoleClaimsRequest{
		RoleId: appRole.RoleID,
		Claims: allClaims,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AppRoleSvc) FetchRole(ctx context.Context, appID string, roleType model.RoleType) (*ClaimsDTO, error) {
	appRole, err := s.roleRepo.FetchRoleByAppIDAndType(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	role, err := s.roleCli.GetRole(ctx, appRole.RoleID)
	if err != nil {
		return nil, err
	}

	if role == nil || role.Role == nil {
		return &ClaimsDTO{}, nil
	}

	allClaims := s.fromProtoClaims(role.Role.Claims)
	nativeClaims, appClaims := s.classifyClaims(allClaims)
	return &ClaimsDTO{NativeClaims: nativeClaims, AppClaims: appClaims}, nil
}

func (s *AppRoleSvc) classifyClaims(claims []*model.Claim) (natives []*model.Claim, apps []*model.Claim) {
	for _, claim := range claims {
		if s.isNativeSvc(claim.Service) {
			natives = append(natives, claim)
		} else {
			apps = append(apps, claim)
		}
	}
	return
}

func (s *AppRoleSvc) isNativeSvc(svcName string) bool {
	for _, nativeSvc := range s.nativeServices {
		if svcName == nativeSvc {
			return true
		}
	}
	return false
}

func (s *AppRoleSvc) DeleteRoles(ctx context.Context, appID string) error {
	roles, err := s.roleRepo.FetchByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	for _, role := range roles {
		if _, err := s.roleCli.DeleteRole(ctx, role.RoleID); err != nil {
			return err
		}
	}

	return s.roleRepo.DeleteByAppID(ctx, appID)
}

func (s *AppRoleSvc) fromProtoClaims(claims []*protomodel.Claim) []*model.Claim {
	ret := make([]*model.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, s.fromProtoClaim(claim))
	}
	return ret
}

func (s *AppRoleSvc) fromProtoClaim(claim *protomodel.Claim) *model.Claim {
	return &model.Claim{
		Service: claim.Service,
		Action:  claim.Action,
	}
}

func (s *AppRoleSvc) nativeClaimsToProto(t model.RoleType, claims []*model.Claim) ([]*protomodel.Claim, error) {
	ret := make([]*protomodel.Claim, 0, len(claims))
	for _, claim := range claims {
		protoClaim, found := s.findMatchingAvailableClaim(t, claim)
		if !found {
			return nil, apierr.Unauthorized(errors.New("contains claim that is not permitted"))
		}
		ret = append(ret, protoClaim)
	}
	return ret, nil
}

func (s *AppRoleSvc) findMatchingAvailableClaim(t model.RoleType, claim *model.Claim) (*protomodel.Claim, bool) {
	claims := s.typeRule[t].AvailableClaims
	for _, protoClaim := range claims {
		if protoClaim.Service == claim.Service && protoClaim.Action == claim.Action {
			return protoClaim, true
		}
	}
	return nil, false
}

func (s *AppRoleSvc) appClaimsToProto(ctx context.Context, claims model.Claims) ([]*protomodel.Claim, error) {
	appIds := make([]string, 0, len(claims))
	for _, claim := range claims {
		appIds = append(appIds, claim.Service)
	}

	apps, err := s.appSvc.ReadAllByAppIDs(ctx, appIds)
	if err != nil {
		return nil, err
	}

	if len(apps) != len(appIds) {
		return nil, apierr.NotFound(errors.New("app not found"))
	}

	ret := make([]*protomodel.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, &protomodel.Claim{
			Service: claim.Service,
			Action:  claim.Action,
			Scope:   []string{"channel-{id}"},
		})
	}

	return ret, nil
}

func (s *AppRoleSvc) GetAvailableNativeClaims(ctx context.Context, roleType model.RoleType) ([]*model.Claim, error) {
	claims := s.typeRule[roleType].AvailableClaims
	ret := make([]*model.Claim, 0, len(claims))
	for _, claim := range claims {
		ret = append(ret, &model.Claim{
			Service: claim.Service,
			Action:  claim.Action,
		})
	}
	return ret, nil
}

func (s *AppRoleSvc) DeleteAppSecret(ctx context.Context, appID string) error {
	return s.secretRepo.Delete(ctx, appID)
}

func (s *AppRoleSvc) RefreshAppSecret(ctx context.Context, appID string) (string, error) {
	token, err := generateSecret()
	if err != nil {
		return "", err
	}

	if err := s.secretRepo.Save(ctx, &model.AppSecret{
		AppID:  appID,
		Secret: token,
	}); err != nil {
		return "", err
	}

	return token, nil
}

func generateSecret() (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	secret := base64.URLEncoding.EncodeToString(randomBytes)
	return secret, nil
}

func (s *AppRoleSvc) HasIssuedBefore(ctx context.Context, appID string) (bool, error) {
	_, err := s.secretRepo.FetchByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return false, nil
	}
	return true, nil
}
