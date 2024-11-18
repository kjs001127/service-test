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

const (
	channelID = "1"
)

type AppIntegrationTestSuite struct {
	suite.Suite

	testHelper *TestHelper

	appLifecycleSvc      svc.AppLifecycleSvc
	appQuerySvc          svc.AppQuerySvc
	appInstallSvc        svc.AppInstallSvc
	installedAppQuerySvc *svc.InstalledAppQuerySvc
}

func (a *AppIntegrationTestSuite) SetupTest() {
	a.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&a.appLifecycleSvc),
		fx.Populate(&a.appQuerySvc),
		fx.Populate(&a.appInstallSvc),
		fx.Populate(&a.installedAppQuerySvc),
	)
	a.testHelper.TruncateAll()
}

func (a *AppIntegrationTestSuite) TearDownSuite() {
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
	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

	app, err := a.appQuerySvc.Read(context.Background(), created.ID)

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().NotEmpty(app.ID)
	a.Require().Equal(app.Title, "test app")
}

func (a *AppIntegrationTestSuite) TestAppUpdate() {
	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

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
	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}
	err = a.appLifecycleSvc.Delete(context.Background(), created.ID)

	a.Require().NoError(err)
}

func (a *AppIntegrationTestSuite) TestReadAllByAppIDs() {
	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

	apps, err := a.appQuerySvc.ReadAllByAppIDs(context.Background(), []string{created.ID})

	a.Require().NoError(err)
	a.Require().NotEmpty(apps)
	a.Require().Len(apps, 1)
	a.Require().Equal(apps[0].Title, "test app")
}

func (a *AppIntegrationTestSuite) TestInstallApp() {
	ctx := context.Background()

	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

	app, err := a.appInstallSvc.InstallApp(ctx, channelID, created)

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().Equal(app.Title, "test app")
}

func (a *AppIntegrationTestSuite) TestInstallAppById() {
	ctx := context.Background()

	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

	installationID := appmodel.InstallationID{
		AppID:     created.ID,
		ChannelID: channelID,
	}

	app, err := a.appInstallSvc.InstallAppById(ctx, installationID)

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().Equal(app.Title, "test app")
}

func (a *AppIntegrationTestSuite) TestUninstallApp() {
	ctx := context.Background()

	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}

	installationID := appmodel.InstallationID{
		AppID:     created.ID,
		ChannelID: channelID,
	}

	_, err = a.appInstallSvc.InstallAppById(ctx, installationID)
	a.Require().NoError(err)

	err = a.appInstallSvc.UnInstallApp(ctx, installationID)
	a.Require().NoError(err)
}

func (a *AppIntegrationTestSuite) TestReadInstalledApp() {

	installed, err := a.createAndInstall()
	a.Require().NoError(err)
	a.Require().NotNil(installed)

	app, err := a.installedAppQuerySvc.QueryInstalledApp(context.Background(), appmodel.InstallationID{
		AppID:     installed.ID,
		ChannelID: channelID,
	})
	a.Require().NoError(err)
	a.Require().Equal(installed.ID, app.ID)
}

func (a *AppIntegrationTestSuite) TestReadInstalledApps() {
	ctx := context.Background()

	installCnt := 10
	var apps []*appmodel.App
	for i := 0; i < installCnt; i++ {
		installed, err := a.createAndInstall()
		a.Require().NoError(err)
		apps = append(apps, installed)
	}

	installedApps, err := a.installedAppQuerySvc.QueryInstalledAppsByChannelID(ctx, channelID)
	a.Require().NoError(err)
	a.Require().Equal(installCnt, len(installedApps))
}

func (a *AppIntegrationTestSuite) TestReadInstalledAppsEmpty() {
	ctx := context.Background()

	installedApps, err := a.installedAppQuerySvc.QueryInstalledAppsByChannelID(ctx, channelID)
	a.Require().NoError(err)
	a.Require().Equal(0, len(installedApps))
}

func (a *AppIntegrationTestSuite) createAndInstall() (*appmodel.App, error) {
	ctx := context.Background()

	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		return nil, err
	}

	_, err = a.appInstallSvc.InstallApp(ctx, channelID, created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func TestAppIntegrationSuite(t *testing.T) {
	suite.Run(t, new(AppIntegrationTestSuite))
}
