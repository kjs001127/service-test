package svc_test

import (
	"context"
	"testing"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	. "github.com/channel-io/ch-app-store/test/integration"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

type AppIntegrationTestSuite struct {
	suite.Suite

	testHelper *TestHelper

	appLifecycleSvc svc.AppLifecycleSvc
	appQuerySvc     svc.AppQuerySvc
	appRepository   svc.AppRepository
	appInstallSvc   *svc.AppInstallSvc
	appInstallRepo  svc.AppInstallationRepository
}

func (a *AppIntegrationTestSuite) SetupTest() {
	a.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&a.appLifecycleSvc),
		fx.Populate(&a.appQuerySvc),
		fx.Populate(&a.appRepository),
		fx.Populate(&a.appInstallSvc),
		fx.Populate(&a.appInstallRepo),
	)
}

func (a *AppIntegrationTestSuite) TearDownSuite() {
	a.testHelper.TruncateAll()
	a.testHelper.Stop()
}

func (a *AppIntegrationTestSuite) TestAppCreate() {
	app, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().NotEmpty(app.ID)
	a.Require().Equal(app.Title, "test app")
}

func (a *AppIntegrationTestSuite) TestAppRead() {
	created, _ := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})

	app, err := a.appQuerySvc.Read(context.Background(), created.ID)

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().NotEmpty(app.ID)
	a.Require().Equal(app.Title, "test app")
}

func (a *AppIntegrationTestSuite) TestAppUpdate() {
	created, _ := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})

	app, err := a.appLifecycleSvc.Update(context.Background(), &appmodel.App{
		ID:    created.ID,
		Title: "new title",
	})

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().NotEmpty(app.ID)
	a.Require().Equal(app.Title, "new title")
}

func (a *AppIntegrationTestSuite) TestAppDelete() {
	created, _ := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})

	err := a.appLifecycleSvc.Delete(context.Background(), created.ID)

	a.Require().NoError(err)
}

func (a *AppIntegrationTestSuite) TestReadPublicApps() {
	_, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title:     "test app",
		IsPrivate: false,
	})

	apps, err := a.appQuerySvc.ReadPublicApps(context.Background(), "0", 500)

	a.Require().NoError(err)
	a.Require().NotEmpty(apps)
}

func (a *AppIntegrationTestSuite) TestReadAllByAppIDs() {
	created, _ := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title:     "test app",
		IsPrivate: false,
	})

	apps, err := a.appQuerySvc.ReadAllByAppIDs(context.Background(), []string{created.ID})

	a.Require().NoError(err)
	a.Require().NotEmpty(apps)
	a.Require().Len(apps, 1)
	a.Require().Equal(apps[0].Title, "test app")
}

func TestAppIntegrationSuite(t *testing.T) {
	suite.Run(t, new(AppIntegrationTestSuite))
}
