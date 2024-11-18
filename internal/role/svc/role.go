package svc

import (
	"context"
	"fmt"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/role/model"
	authgen "github.com/channel-io/ch-app-store/internal/shared/general"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/internal/shared/principal/front"
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type ClaimsRequestWithID struct {
	ID           string       `json:"id"`
	AppID        string       `json:"appId"`
	NativeClaims model.Claims `json:"nativeClaims,omitempty"`
	AppClaims    model.Claims `json:"appClaims,omitempty"`
}

type ClaimsRequest struct {
	AppID        string         `json:"appId"`
	Type         model.RoleType `json:"type"`
	NativeClaims model.Claims   `json:"nativeClaims,omitempty"`
	AppClaims    model.Claims   `json:"appClaims,omitempty"`
}

type ClaimsResponse struct {
	ID           string         `json:"id"`
	AppID        string         `json:"appId"`
	Type         model.RoleType `json:"type"`
	NativeClaims model.Claims   `json:"nativeClaims,omitempty"`
	AppClaims    model.Claims   `json:"appClaims,omitempty"`
}

type AppRoleID struct {
	AppRoleID string
	AppID     string
}

type ClaimManager interface {
	DefaultClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error)
	AvailableClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error)
}

type AppRoleSvc struct {
	roleCli            authgen.RoleFetcher
	roleRepo           AppRoleRepository
	nativeClaimManager ClaimManager
	querySvc           svc.AppQuerySvc
}

func NewAppRoleSvc(
	roleCli authgen.RoleFetcher,
	roleRepo AppRoleRepository,
	nativeClaimManager ClaimManager,
	querySvc svc.AppQuerySvc,
) *AppRoleSvc {
	return &AppRoleSvc{
		roleCli:            roleCli,
		roleRepo:           roleRepo,
		nativeClaimManager: nativeClaimManager,
		querySvc:           querySvc,
	}
}

func (s *AppRoleSvc) CreateDefaultRoles(ctx context.Context, appID string) error {
	for _, roleType := range model.AvailableRoleTypes {
		if _, err := s.CreateRole(ctx, &ClaimsRequest{Type: roleType, AppID: appID}); err != nil {
			return err
		}
	}
	return nil
}

func (s *AppRoleSvc) CreateRole(ctx context.Context, request *ClaimsRequest) (*ClaimsResponse, error) {

	claims, err := s.checkAndMergeClaims(ctx, request)
	if err != nil {
		return nil, err
	}

	defaults, err := s.defaultClaims(ctx, request.AppID, request.Type)
	if err != nil {
		return nil, err
	}

	claims = append(claims, defaults...)

	res, err := s.roleCli.CreateRole(ctx, &service.CreateRoleRequest{
		Claims:                toProtoClaims(request.Type, request.AppID, claims),
		AllowedPrincipalTypes: principalTypes(request.Type),
		AllowedGrantTypes:     grantTypes(request.Type),
	})
	if err != nil {
		return nil, err
	}

	latestVersion := 0
	if latest, err := s.roleRepo.FindLatestRole(ctx, request.AppID, request.Type); err == nil {
		latestVersion = latest.Version
	}

	saved, err := s.roleRepo.Save(ctx, &model.AppRole{
		AppID:   request.AppID,
		Type:    request.Type,
		Version: latestVersion + 1,

		RoleID: res.Role.Id,
		Credentials: &model.Credentials{
			ClientSecret: res.Credentials.ClientSecret,
			ClientID:     res.Credentials.ClientId,
		},
	})
	if err != nil {
		return nil, err
	}

	return &ClaimsResponse{
		ID:           saved.ID,
		Type:         request.Type,
		AppID:        request.AppID,
		NativeClaims: request.NativeClaims,
		AppClaims:    request.AppClaims,
	}, nil
}

func (s *AppRoleSvc) checkAndMergeClaims(ctx context.Context, set *ClaimsRequest) (model.Claims, error) {
	availableNativeClaims, err := s.nativeClaimManager.AvailableClaims(ctx, set.AppID, set.Type)
	if err != nil {
		return nil, err
	}

	if denied := s.sub(set.NativeClaims, availableNativeClaims); len(denied) > 0 {
		return nil, apierr.Unauthorized(errors.New("claim denied"))
	}

	var allClaims []*model.Claim
	allClaims = append(allClaims, set.NativeClaims...)
	allClaims = append(allClaims, set.AppClaims...)
	return allClaims, nil
}

func (s *AppRoleSvc) defaultClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error) {
	defaultNativeClaims, err := s.nativeClaimManager.DefaultClaims(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	defaultAppClaims := s.appDefaultClaim(appID)

	ret := make(model.Claims, 0, len(defaultAppClaims)+len(defaultNativeClaims))
	ret = append(ret, defaultNativeClaims...)
	ret = append(ret, defaultAppClaims...)

	return ret, nil
}

func (s *AppRoleSvc) pickAppClaims(ctx context.Context, claims model.Claims) (model.Claims, error) {
	var appIDs []string
	for _, claim := range claims {
		appIDs = append(appIDs, claim.Service)
	}

	apps, err := s.querySvc.ReadAllByAppIDs(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	var ret model.Claims
	for _, app := range apps {
		ret = append(ret, &model.Claim{Service: app.ID, Action: "*"})
	}

	return ret, nil
}

func (s *AppRoleSvc) pickNativeClaims(ctx context.Context, appID string, roleType model.RoleType, claims model.Claims) (model.Claims, error) {
	availableNativeClaims, err := s.nativeClaimManager.AvailableClaims(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	var ret model.Claims
	for _, claim := range claims {
		if s.contains(claim, availableNativeClaims) {
			ret = append(ret, claim)
		}
	}

	return ret, nil
}

func (s *AppRoleSvc) appDefaultClaim(appID string) model.Claims {
	return model.Claims{{Service: appID, Action: "*"}}
}

func toProtoClaims(t model.RoleType, appID string, claims model.Claims) []*protomodel.Claim {
	var ret []*protomodel.Claim
	for _, claim := range claims {
		ret = append(ret, &protomodel.Claim{
			Service: claim.Service,
			Action:  claim.Action,
			Scope:   scope(t, appID),
		})
	}

	return ret
}

func scope(t model.RoleType, appID string) []string {
	switch t {
	case model.RoleTypeChannel:
		return []string{fmt.Sprintf("app-%s", appID), "channel-{id}"}
	case model.RoleTypeUser:
		return []string{fmt.Sprintf("app-%s", appID), "user-{id}", "channel-{id}"}
	case model.RoleTypeManager:
		return []string{fmt.Sprintf("app-%s", appID), "manager-{id}", "channel-{id}"}
	case model.RoleTypeApp:
		fallthrough
	default:
		return []string{fmt.Sprintf("app-%s", appID)}
	}
}

func principalTypes(t model.RoleType) []string {
	switch t {
	case model.RoleTypeUser:
		return []string{front.XSessionHeader}
	case model.RoleTypeManager:
		return []string{desk.XAccountHeader}
	}
	return nil
}

func grantTypes(t model.RoleType) []protomodel.GrantType {
	switch t {
	case model.RoleTypeApp:
		fallthrough
	case model.RoleTypeChannel:
		return []protomodel.GrantType{protomodel.GrantType_GRANT_TYPE_CLIENT_CREDENTIALS, protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN}

	case model.RoleTypeUser:
		fallthrough
	case model.RoleTypeManager:
		return []protomodel.GrantType{protomodel.GrantType_GRANT_TYPE_PRINCIPAL, protomodel.GrantType_GRANT_TYPE_REFRESH_TOKEN}
	}

	return nil
}

func (s *AppRoleSvc) UpdateRole(ctx context.Context, request *ClaimsRequestWithID) error {
	appRole, err := s.roleRepo.Find(ctx, request.ID)
	if err != nil {
		return err
	}

	claims, err := s.checkAndMergeClaims(ctx, &ClaimsRequest{
		AppClaims:    request.AppClaims,
		NativeClaims: request.NativeClaims,
		Type:         appRole.Type,
		AppID:        appRole.AppID,
	})
	if err != nil {
		return err
	}

	defaults, err := s.defaultClaims(ctx, appRole.AppID, appRole.Type)
	if err != nil {
		return err
	}

	claims = append(claims, defaults...)

	_, err = s.roleCli.UpdateRole(ctx, &service.ReplaceRoleClaimsRequest{
		RoleId: appRole.RoleID,
		Claims: toProtoClaims(appRole.Type, appRole.AppID, claims),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AppRoleSvc) FetchAvailableNativeClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error) {
	return s.nativeClaimManager.AvailableClaims(ctx, appID, roleType)
}

func (s *AppRoleSvc) AppendClaimsToRole(ctx context.Context, request *ClaimsRequestWithID) error {
	appRole, err := s.roleRepo.Find(ctx, request.ID)
	if err != nil {
		return err
	}

	claimsToAdd, err := s.checkAndMergeClaims(ctx, &ClaimsRequest{
		AppClaims:    request.AppClaims,
		NativeClaims: request.NativeClaims,
		Type:         appRole.Type,
		AppID:        appRole.AppID,
	})
	if err != nil {
		return err
	}

	role, err := s.roleCli.GetRole(ctx, appRole.RoleID)
	if err != nil {
		return err
	}
	if role == nil || role.Role == nil {
		return errors.New("role does not exist")
	}

	allClaims := append(role.Role.Claims, toProtoClaims(appRole.Type, appRole.AppID, claimsToAdd)...)

	_, err = s.roleCli.UpdateRole(ctx, &service.ReplaceRoleClaimsRequest{
		RoleId: appRole.RoleID,
		Claims: allClaims,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *AppRoleSvc) FetchClaims(ctx context.Context, appRole *model.AppRole) (*ClaimsResponse, error) {
	role, err := s.roleCli.GetRole(ctx, appRole.RoleID)
	if err != nil {
		return nil, err
	}

	if role == nil || role.Role == nil {
		return nil, apierr.NotFound(errors.New("role cli response is emtpy"))
	}

	allClaims := s.fromProtoClaims(role.Role.Claims)
	return s.classifyClaims(ctx, appRole, allClaims)
}

func (s *AppRoleSvc) FetchLatestRole(ctx context.Context, appID string, roleType model.RoleType) (*ClaimsResponse, error) {
	appRole, err := s.roleRepo.FindLatestRole(ctx, appID, roleType)
	if err != nil {
		return nil, err
	}

	return s.FetchClaims(ctx, appRole)
}

func (s *AppRoleSvc) FetchRole(ctx context.Context, id string) (*ClaimsResponse, error) {
	appRole, err := s.roleRepo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.FetchClaims(ctx, appRole)
}

func (s *AppRoleSvc) roleWithMaxVersion(roles []*model.AppRole) *model.AppRole {
	var ret *model.AppRole
	for _, role := range roles {
		if ret == nil || ret.Version < role.Version {
			ret = role
		}
	}
	return ret
}

func (s *AppRoleSvc) classifyClaims(ctx context.Context, role *model.AppRole, claims model.Claims) (*ClaimsResponse, error) {
	nativeCh := make(chan model.Claims)
	appCh := make(chan model.Claims)

	go func() {
		natives, err := s.pickNativeClaims(ctx, role.AppID, role.Type, claims)
		if err != nil {
			close(nativeCh)
		}

		nativeDefaults, err := s.nativeClaimManager.DefaultClaims(ctx, role.AppID, role.Type)
		if err != nil {
			close(nativeCh)
		}

		nativeCh <- s.sub(natives, nativeDefaults)
	}()

	go func() {
		apps, err := s.pickAppClaims(ctx, claims)
		if err != nil {
			close(appCh)
		}

		appDefaults := s.appDefaultClaim(role.AppID)
		appCh <- s.sub(apps, appDefaults)
	}()

	apps, ok := <-appCh
	if !ok {
		return nil, errors.New("app claim fetch failed")
	}

	natives, ok := <-nativeCh
	if !ok {
		return nil, errors.New("native claim fetch failed")
	}

	return &ClaimsResponse{
		Type:         role.Type,
		AppID:        role.AppID,
		NativeClaims: natives,
		AppClaims:    apps,
	}, nil
}

func (s *AppRoleSvc) sub(claims model.Claims, excludes model.Claims) model.Claims {
	var ret model.Claims
	for _, claim := range claims {
		if !s.contains(claim, excludes) {
			ret = append(ret, claim)
		}
	}

	return ret
}

func (s *AppRoleSvc) contains(check *model.Claim, targets model.Claims) bool {
	for _, target := range targets {
		if target.Equal(check) {
			return true
		}
	}
	return false
}

func (s *AppRoleSvc) DeleteRoles(ctx context.Context, appID string) error {
	roles, err := s.roleRepo.FindAllByAppID(ctx, appID)
	if apierr.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	if err := s.roleRepo.DeleteByAppID(ctx, appID); err != nil {
		return err
	}

	for _, role := range roles {
		if _, err := s.roleCli.DeleteRole(ctx, role.RoleID); err != nil {
			return err
		}
	}

	return nil
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

type ClaimsFactory func(string) model.Claims

type StaticClaimManager map[model.RoleType]Claims

type Claims struct {
	DefaultClaimsOf   ClaimsFactory
	AvailableClaimsOf ClaimsFactory
}

func (s StaticClaimManager) DefaultClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error) {
	return s[roleType].DefaultClaimsOf(appID), nil
}

func (s StaticClaimManager) AvailableClaims(ctx context.Context, appID string, roleType model.RoleType) (model.Claims, error) {
	return s[roleType].AvailableClaimsOf(appID), nil
}
