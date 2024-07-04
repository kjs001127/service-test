package dao_test

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
	appID     = "1"
	channelID = "1"
)

type AppDAOTestSuite struct {
	suite.Suite

	testApp *TestHelper

	appRepository             svc.AppRepository
	appInstallationRepository svc.AppInstallationRepository
}

func (a *AppDAOTestSuite) SetupTest() {
	a.testApp = NewTestHelper(
		testOpts,
		fx.Populate(&a.appRepository),
		fx.Populate(&a.appInstallationRepository),
	)
	a.testApp.TruncateAll()
}

func (a *AppDAOTestSuite) TearDownSuite() {
	a.testApp.Stop()
}

func (a *AppDAOTestSuite) TestAppSave() {
	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	res, err := a.appRepository.Save(ctx, app)

	a.Require().NoError(err)
	a.Require().NotNil(res)

	res, err = a.appRepository.FindApp(ctx, appID)

	a.Require().NoError(err)
	a.Require().NotNil(res)
}

func (a *AppDAOTestSuite) TestAppFind() {
	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	res, err := a.appRepository.FindApp(ctx, appID)

	a.Require().NoError(err)
	a.Require().NotNil(res)
}

func (a *AppDAOTestSuite) TestAppDelete() {
	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	err := a.appRepository.Delete(ctx, appID)

	a.Require().NoError(err)
}

func (a *AppDAOTestSuite) TestAppInstallationSave() {
	appChannel := &appmodel.AppInstallation{
		ChannelID: channelID,
		AppID:     appID,
	}

	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	err := a.appInstallationRepository.Save(ctx, appChannel)

	a.Require().NoError(err)
}

func (a *AppDAOTestSuite) TestAppInstallationDelete() {
	appChannel := &appmodel.AppInstallation{
		ChannelID: channelID,
		AppID:     appID,
	}

	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	_ = a.appInstallationRepository.Save(ctx, appChannel)
	err := a.appInstallationRepository.DeleteByAppID(ctx, appID)

	a.Require().NoError(err)
}

func (a *AppDAOTestSuite) TestAppInstallationFind() {
	appChannel := &appmodel.AppInstallation{
		ChannelID: channelID,
		AppID:     appID,
	}

	installationID := appmodel.InstallationID{
		ChannelID: channelID,
		AppID:     appID,
	}

	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, err := a.appRepository.Save(ctx, app)
	a.Require().NoError(err)

	err = a.appInstallationRepository.Save(ctx, appChannel)
	a.Require().NoError(err)

	res, err := a.appInstallationRepository.Fetch(ctx, installationID)

	a.Require().NoError(err)
	a.Require().NotNil(res)
}

func TestAppDAOs(t *testing.T) {
	suite.Run(t, new(AppDAOTestSuite))
}
