package svc

import (
	"context"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/internal/shared/principal/desk"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

type ManagerChannelAgreementSvc struct {
	svc *ChannelAgreementSvc
}

func NewManagerChannelAgreementSvc(svc *ChannelAgreementSvc) *ManagerChannelAgreementSvc {
	return &ManagerChannelAgreementSvc{svc: svc}
}

func (m *ManagerChannelAgreementSvc) Agree(ctx context.Context, manager desk.Manager, channelID string, appRoleIDs []string) error {
	role, err := manager.Role(ctx)
	if err != nil {
		return err
	}

	if !role.IsOwner() {
		return err
	}

	return m.svc.Agree(ctx, channelID, appRoleIDs)
}

type ChannelAgreementSvc struct {
	agreementRepo ChannelRoleAgreementRepo
	roleSvc       *AppRoleSvc
}

func NewChannelAgreementSvc(repo ChannelRoleAgreementRepo, roleSvc *AppRoleSvc) *ChannelAgreementSvc {
	return &ChannelAgreementSvc{agreementRepo: repo, roleSvc: roleSvc}
}

func (a *ChannelAgreementSvc) Agree(ctx context.Context, channelID string, appRoleIDs []string) error {
	return tx.Do(ctx, func(ctx context.Context) error {
		for _, roleID := range appRoleIDs {
			if _, err := a.roleSvc.FetchRole(ctx, roleID); err != nil {
				return err
			}
			agreement := &model.ChannelRoleAgreement{
				ChannelID: channelID,
				AppRoleID: roleID,
			}
			if err := a.agreementRepo.Save(ctx, agreement); err != nil {
				return err
			}
		}
		return nil
	})
}

var roleTypesForAgree = []model.RoleType{
	model.RoleTypeChannel,
	model.RoleTypeManager,
	model.RoleTypeUser,
}

func (a *ChannelAgreementSvc) FetchUnAgreedRoles(ctx context.Context, id appmodel.InstallationID) ([]*ClaimsResponse, error) {
	roles, err := a.agreementRepo.FindLatestUnAgreedRoles(ctx, id.ChannelID, []string{id.AppID}, roleTypesForAgree)
	if err != nil {
		return nil, err
	}

	ret := make([]*ClaimsResponse, 0, len(roles))
	for _, role := range roles {
		resp, err := a.roleSvc.FetchClaims(ctx, role)
		if err != nil {
			return nil, err
		}
		ret = append(ret, resp)
	}

	return ret, nil
}

type AgreementBulkFetchResponse struct {
	AppID     string
	AppRoleID string
}

func (a *ChannelAgreementSvc) BulkFetchUnAgreedRoles(ctx context.Context, channelID string, appIDs []string) ([]AgreementBulkFetchResponse, error) {

	roles, err := a.agreementRepo.FindLatestUnAgreedRoles(ctx, channelID, appIDs, roleTypesForAgree)
	if err != nil {
		return nil, err
	}

	ret := make([]AgreementBulkFetchResponse, 0, len(roles))
	for _, role := range roles {
		ret = append(ret, AgreementBulkFetchResponse{
			AppID:     role.AppID,
			AppRoleID: role.ID,
		})
	}

	return ret, nil
}
