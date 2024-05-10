package permission_test

import (
	"context"
	"testing"

	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/managerfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/permissionfx"
	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	crudSvc "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	managersvc "github.com/channel-io/ch-app-store/internal/manager/svc"
	permission "github.com/channel-io/ch-app-store/internal/permission/svc"
	. "github.com/channel-io/ch-app-store/test/integration"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
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
	helper *TestHelper

	appSvc             permission.AccountAppPermissionSvc
	installSvc         *managersvc.ManagerAwareInstallSvc
	lifecycleSvc       crudSvc.AppLifecycleSvc
	managerRoleFetcher mockaccount.ManagerRoleFetcher
}

var suite PermissionTestSuite

var _ = BeforeSuite(func() {
	suite.helper = NewTestHelper(
		configfx.Values,
		datadogfx.Datadog,
		permissionfx.Permission,
		appfx.App,
		managerfx.Manager,

		fx.Populate(&suite.appSvc),
		fx.Populate(&suite.installSvc),
		fx.Populate(&suite.lifecycleSvc),
		Mock[account.ManagerRoleFetcher](&suite.managerRoleFetcher),
	)
})

var _ = BeforeEach(func() {
	suite.managerRoleFetcher = mockaccount.ManagerRoleFetcher{}
})

var _ = AfterSuite(func() {
	suite.helper.Stop()
	suite.helper.WithPreparedTables("app_accounts")
})

var _ = Describe("CreateApp", func() {
	Context("when creating app", func() {
		It("should create app", func() {
			ctx := context.Background()

			app, err := suite.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())
			Expect(app.Title).To(Equal(testTitle))
		})
	})
})

var _ = Describe("DeleteApp", func() {
	Context("when deleting app", func() {
		It("should delete app", func() {
			ctx := context.Background()

			app, err := suite.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())

			err = suite.appSvc.DeleteApp(ctx, app.ID, ownerAccountID)
			Expect(err).To(BeNil())
		})
	})
})

var _ = Describe("InstallApp", func() {
	Context("when install private App by owner", func() {
		It("should install app", func() {
			ctx := context.Background()

			manager := account.Manager{
				ID:        ownerManagerID,
				RoleID:    ownerRoleID,
				AccountID: ownerAccountID,
			}

			managerRole := account.ManagerRole{
				ID:          ownerRoleID,
				RoleType:    ownerType,
				Permissions: []account.Permission{{Action: channelPermission}},
			}
			suite.managerRoleFetcher.EXPECT().FetchRole(mock.Anything, ownerRoleID).Return(managerRole, nil)

			app, err := suite.appSvc.CreateApp(ctx, testTitle, ownerAccountID)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())

			installationID := appmodel.InstallationID{
				AppID: app.ID,
			}

			installedApp, err := suite.installSvc.Install(ctx, manager, installationID)
			Expect(err).To(BeNil())
			Expect(installedApp).ToNot(BeNil())
			Expect(installedApp.ID).To(Equal(app.ID))
		})
	})

	Context("when install private App by non-owner", func() {
		It("should return error", func() {
			ctx := context.Background()

			manager := account.Manager{
				ID:        nonOwnerRoleID,
				RoleID:    nonOwnerRoleID,
				AccountID: nonOwnerAccountID,
			}

			managerRole := account.ManagerRole{
				ID:          nonOwnerRoleID,
				RoleType:    nonOwnerType,
				Permissions: []account.Permission{{Action: channelPermission}},
			}

			suite.managerRoleFetcher.EXPECT().FetchRole(mock.Anything, nonOwnerRoleID).Return(managerRole, nil)
			app, _ := suite.lifecycleSvc.Create(ctx, &appmodel.App{Title: testTitle, IsPrivate: true})

			installationID := appmodel.InstallationID{
				AppID:     app.ID,
				ChannelID: channelID,
			}
			app, err := suite.installSvc.Install(ctx, manager, installationID)

			Expect(err).To(Not(BeNil()))
			Expect(app).To(BeNil())
		})
	})
})

func TestPermissionSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PermissionTestSuite")
}
