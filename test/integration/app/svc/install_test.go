package svc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"

	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	. "github.com/channel-io/ch-app-store/test/integration"
)

type InstallTestSuite struct {
	suite.Suite

	testHelper *TestHelper

	appLifecycleSvc      svc.AppLifecycleSvc
	appInstallSvc        svc.AppInstallSvc
	installedAppQuerySvc *svc.InstalledAppQuerySvc
}

func (a *InstallTestSuite) SetupTest() {
	a.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&a.appLifecycleSvc),
		fx.Populate(&a.appInstallSvc),
		fx.Populate(&a.installedAppQuerySvc),
	)
	a.testHelper.TruncateAll()
}

func (a *InstallTestSuite) TearDownSuite() {
	a.testHelper.Stop()
}

func (a *InstallTestSuite) TestInstallApp() {
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

func (a *InstallTestSuite) TestInstallAppById() {
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

func (a *InstallTestSuite) TestUninstallApp() {
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

func (a *InstallTestSuite) TestReadInstalledApp() {

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

func (a *InstallTestSuite) TestReadInstalledApps() {
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

func (a *InstallTestSuite) TestReadInstalledAppsEmpty() {
	ctx := context.Background()

	installedApps, err := a.installedAppQuerySvc.QueryInstalledAppsByChannelID(ctx, channelID)
	a.Require().NoError(err)
	a.Require().Equal(0, len(installedApps))
}

func (a *InstallTestSuite) createAndInstall() (*appmodel.App, error) {
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

func TestAppInstall(t *testing.T) {
	suite.Run(t, new(InstallTestSuite))
}
