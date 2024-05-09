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
	testApp *TestApp

	permission.AccountAppPermissionSvc
	*managersvc.ManagerAwareInstallSvc
	permission.AppAccountRepo
	crudSvc.AppCrudSvc
	managerRoleFetcher mockaccount.ManagerRoleFetcher
}

var suite PermissionTestSuite

var _ = BeforeSuite(func() {
	suite.testApp = NewTestApp(
		Populate(&suite.AccountAppPermissionSvc),
		Populate(&suite.ManagerAwareInstallSvc),
		Populate(&suite.AppAccountRepo),
		Populate(&suite.AppCrudSvc),
		Mock[account.ManagerRoleFetcher](&suite.managerRoleFetcher),
	)
})

var _ = BeforeEach(func() {
	suite.managerRoleFetcher = mockaccount.ManagerRoleFetcher{}
})

var _ = AfterSuite(func() {
	suite.testApp.Stop()
	suite.testApp.WithPreparedTables("app_accounts")
})

var _ = Describe("CreateApp", func() {
	Context("when creating app", func() {
		It("should create app", func() {
			ctx := context.Background()

			app, err := suite.CreateApp(ctx, testTitle, ownerAccountID)
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

			app, err := suite.CreateApp(ctx, testTitle, ownerAccountID)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())

			err = suite.DeleteApp(ctx, app.ID, ownerAccountID)
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

			app, err := suite.CreateApp(ctx, testTitle, ownerAccountID)
			Expect(err).To(BeNil())
			Expect(app).ToNot(BeNil())

			installationID := appmodel.InstallationID{
				AppID: app.ID,
			}

			installedApp, err := suite.Install(ctx, manager, installationID)
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
			app, _ := suite.AppCrudSvc.Create(ctx, &appmodel.App{Title: testTitle, IsPrivate: true})

			installationID := appmodel.InstallationID{
				AppID:     app.ID,
				ChannelID: channelID,
			}
			app, err := suite.Install(ctx, manager, installationID)

			Expect(err).To(Not(BeNil()))
			Expect(app).To(BeNil())
		})
	})
})

func TestPermissionSvc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PermissionTestSuite")
}
