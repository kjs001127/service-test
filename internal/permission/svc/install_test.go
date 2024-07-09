package svc_test

import (
	"context"
	mocksvc "github.com/channel-io/ch-app-store/generated/mock/appdisplay/svc"
	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	displaysvc "github.com/channel-io/ch-app-store/generated/mock/permission/svc"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/appdisplay/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/permission/svc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type InstallPermissionSuite struct {
	suite.Suite

	permissionUtil svc.PermissionUtil
	appDisplayRepo *mocksvc.AppDisplayRepository
	appAccountRepo *displaysvc.AppAccountRepo
	roleFetcher    *mockaccount.ManagerRoleFetcher

	svc *svc.ManagerInstallPermissionSvcImpl
}

func (p *InstallPermissionSuite) SetupTest() {
	p.appDisplayRepo = mocksvc.NewAppDisplayRepository(p.T())
	p.appAccountRepo = displaysvc.NewAppAccountRepo(p.T())
	p.roleFetcher = mockaccount.NewManagerRoleFetcher(p.T())
	p.permissionUtil = svc.NewPermissionUtil(p.roleFetcher)

	p.svc = svc.NewManagerInstallPermissionSvc(p.appDisplayRepo, p.permissionUtil, p.appAccountRepo)
}

func (p *InstallPermissionSuite) TestOnInstall() {
	privateAppID := "1"
	publicAppID := "2"
	channelID := "1"

	ownerAccountID := "1"
	notOwnerAccountID := "2"
	generalSettingsAccountID := "999231"
	noPermissionAccountID := "absent"

	ownerRoleID := "1"
	generalSettingsRoleID := "999231"
	noPermissionRoleID := "absent"

	ownerRole := account.ManagerRole{
		ID:       ownerRoleID,
		RoleType: "owner",
	}

	generalSettingsRole := account.ManagerRole{
		ID:       generalSettingsRoleID,
		RoleType: "general_settings",
		Permissions: []account.Permission{
			{
				Action: "generalSettings",
			},
		},
	}

	p.T().Run("private app, manager is owner and developer", func(t *testing.T) {
		p.appDisplayRepo.EXPECT().FindDisplay(mock.Anything, privateAppID).Return(&model.AppDisplay{IsPrivate: true}, nil)
		p.appAccountRepo.EXPECT().Fetch(mock.Anything, privateAppID, ownerAccountID).Return(nil, nil)
		p.roleFetcher.EXPECT().FetchRole(mock.Anything, channelID, ownerRoleID).Return(ownerRole, nil)

		ctx := context.Background()
		manager := account.Manager{
			RoleID:    ownerRoleID,
			ChannelID: channelID,
			AccountID: ownerAccountID,
		}

		installationID := appmodel.InstallationID{
			AppID:     privateAppID,
			ChannelID: channelID,
		}

		err := p.svc.OnInstall(ctx, manager, installationID)

		assert.Nil(t, err)
	})

	p.T().Run("public app, manager has general_settings permission", func(t *testing.T) {
		p.appDisplayRepo.EXPECT().FindDisplay(mock.Anything, publicAppID).Return(&model.AppDisplay{IsPrivate: false}, nil)
		p.roleFetcher.EXPECT().FetchRole(mock.Anything, channelID, generalSettingsAccountID).Return(generalSettingsRole, nil)

		ctx := context.Background()
		manager := account.Manager{
			RoleID:    generalSettingsRoleID,
			ChannelID: channelID,
			AccountID: generalSettingsAccountID,
		}

		installationID := appmodel.InstallationID{
			AppID:     publicAppID,
			ChannelID: channelID,
		}

		err := p.svc.OnInstall(ctx, manager, installationID)

		assert.Nil(t, err)
	})

	p.T().Run("public app, manager does not have general_settings permission", func(t *testing.T) {
		p.appDisplayRepo.EXPECT().FindDisplay(mock.Anything, publicAppID).Return(&model.AppDisplay{IsPrivate: false}, nil)
		p.roleFetcher.EXPECT().FetchRole(mock.Anything, channelID, noPermissionAccountID).Return(account.ManagerRole{}, nil)

		ctx := context.Background()
		manager := account.Manager{
			RoleID:    noPermissionRoleID,
			ChannelID: channelID,
			AccountID: noPermissionAccountID,
		}

		installationID := appmodel.InstallationID{
			AppID:     publicAppID,
			ChannelID: channelID,
		}

		err := p.svc.OnInstall(ctx, manager, installationID)

		assert.NotNil(t, err)
	})

	p.T().Run("private app, manager is not owner", func(t *testing.T) {
		p.appDisplayRepo.EXPECT().FindDisplay(mock.Anything, privateAppID).Return(&model.AppDisplay{IsPrivate: true}, nil)
		p.appAccountRepo.EXPECT().Fetch(mock.Anything, privateAppID, notOwnerAccountID).Return(nil, nil)
		p.roleFetcher.EXPECT().FetchRole(mock.Anything, channelID, noPermissionRoleID).Return(account.ManagerRole{}, nil)

		ctx := context.Background()
		manager := account.Manager{
			RoleID:    noPermissionRoleID,
			ChannelID: channelID,
			AccountID: notOwnerAccountID,
		}

		installationID := appmodel.InstallationID{
			AppID:     privateAppID,
			ChannelID: channelID,
		}

		err := p.svc.OnInstall(ctx, manager, installationID)

		assert.NotNil(t, err)
	})
}

func TestInstallPermissionSuite(t *testing.T) {
	suite.Run(t, new(InstallPermissionSuite))
}
