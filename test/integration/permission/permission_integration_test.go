package permission_test

import (
	"context"
	"testing"

	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	crudSvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	managersvc "github.com/channel-io/ch-app-store/internal/manager/svc"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
	. "github.com/channel-io/ch-app-store/test/integration"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

const (
	ownerManagerID    = "123"
	ownerRoleID       = "843"
	nonOwnerRoleID    = "844"
	testTitle         = "test"
	ownerAccountID    = "10"
	nonOwnerAccountID = "20"
	ownerType         = "owner"
	nonOwnerType      = "non-owner"
	channelPermission = "generalSettings"
	channelID         = "1"
)

type PermissionTestSuite struct {
	suite.Suite

	helper *TestHelper

	appSvc             permission.AccountAppPermissionSvc
	installSvc         *managersvc.ManagerAwareInstallSvc
	lifecycleSvc       crudSvc.AppLifecycleSvc
	managerRoleFetcher mockaccount.ManagerRoleFetcher
}

func (p *PermissionTestSuite) SetupTest() {
	p.helper = NewTestHelper(
		testOpts,
		fx.Populate(&p.appSvc),
		fx.Populate(&p.installSvc),
		fx.Populate(&p.lifecycleSvc),
		Mock[account.ManagerRoleFetcher](&p.managerRoleFetcher),
	)
	p.helper.TruncateAll()
}

func (p *PermissionTestSuite) TearDownSuite() {
	p.helper.Stop()
}

func (p *PermissionTestSuite) TestAppCreate() {
	app, err := p.appSvc.CreateApp(context.Background(), testTitle, ownerAccountID)

	p.Require().NoError(err)
	p.Require().NotNil(app)
	p.Require().NotEmpty(app.ID)
	p.Require().Equal(app.Title, testTitle)
}

func (p *PermissionTestSuite) TestDeleteApp() {
	ctx := context.Background()

	app, err := p.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
	p.Require().NotNil(app)
	p.Require().NoError(err)

	err = p.appSvc.DeleteApp(ctx, app.ID, ownerAccountID)

	p.Require().NoError(err)
}

func (p *PermissionTestSuite) TestInstallPrivateAppByOwner() {
	ctx := context.Background()

	manager := account.Manager{
		ID:        ownerManagerID,
		RoleID:    ownerRoleID,
		AccountID: ownerAccountID,
		ChannelID: channelID,
	}

	managerRole := account.ManagerRole{
		ID:          ownerRoleID,
		RoleType:    ownerType,
		Permissions: []account.Permission{{Action: channelPermission}},
	}
	p.managerRoleFetcher.EXPECT().FetchRole(mock.Anything, channelID, ownerRoleID).Return(managerRole, nil)

	app, err := p.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
	p.Require().NotNil(app)
	p.Require().NoError(err)

	installationID := appmodel.InstallationID{
		AppID: app.ID,
	}

	installedApp, err := p.installSvc.Install(ctx, manager, installationID)
	p.Require().NotNil(installedApp)
	p.Require().NoError(err)
	p.Require().Equal(installedApp.ID, app.ID)
}

func (p *PermissionTestSuite) TestInstallPrivateAppByNonOwner() {
	ctx := context.Background()

	manager := account.Manager{
		ID:        ownerManagerID,
		RoleID:    nonOwnerRoleID,
		AccountID: nonOwnerAccountID,
		ChannelID: channelID,
	}

	managerRole := account.ManagerRole{
		ID:          nonOwnerRoleID,
		RoleType:    nonOwnerType,
		Permissions: []account.Permission{{Action: channelPermission}},
	}
	p.managerRoleFetcher.EXPECT().FetchRole(mock.Anything, channelID, nonOwnerRoleID).Return(managerRole, nil)

	app, err := p.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
	p.Require().NotNil(app)
	p.Require().NoError(err)

	installationID := appmodel.InstallationID{
		AppID: app.ID,
	}

	installedApp, err := p.installSvc.Install(ctx, manager, installationID)
	p.Require().Nil(installedApp)
	p.Require().Error(err)
}

func TestPermissionSvc(t *testing.T) {
	suite.Run(t, new(PermissionTestSuite))
}
