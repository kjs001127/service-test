package dao_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/volatiletech/null/v8"

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
	displayRepo               svc.AppDisplayRepository
}

func (a *AppDAOTestSuite) SetupTest() {
	a.testApp = NewTestHelper(
		testOpts,
		fx.Populate(&a.appRepository),
		fx.Populate(&a.appInstallationRepository),
		fx.Populate(&a.displayRepo),
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
}

func (a *AppDAOTestSuite) TestAppFind() {
	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)

	res, err := a.appRepository.Find(ctx, appID)
	a.Require().NoError(err)
	a.Require().NotNil(res)
	a.Require().Equal(app.ID, res.ID)
}

func (a *AppDAOTestSuite) TestAppDelete() {
	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	err := a.appRepository.Delete(ctx, appID)
	a.Require().NoError(err)

	_, err = a.appRepository.Find(ctx, appID)
	a.Require().True(apierr.IsNotFound(err))
}

func (a *AppDAOTestSuite) TestAppFindAll() {

	var appIDs []string
	cnt := 10
	for i := 0; i < cnt; i++ {
		newApp := &appmodel.App{
			ID: "app-" + strconv.Itoa(i),
		}

		saved, err := a.appRepository.Save(context.Background(), newApp)
		a.Require().NoError(err)

		appIDs = append(appIDs, saved.ID)
	}

	res, err := a.appRepository.FindAll(context.Background(), appIDs)
	a.Require().NoError(err)

	a.Require().Equal(len(appIDs), len(res))
}

func (a *AppDAOTestSuite) TestListPublicApps() {

	var appIDs []string
	cnt := 10
	for i := 0; i < cnt; i++ {
		newApp := &appmodel.App{
			ID:        "app-" + strconv.Itoa(i),
			IsPrivate: false,
		}

		saved, err := a.appRepository.Save(context.Background(), newApp)
		a.Require().NoError(err)

		appIDs = append(appIDs, saved.ID)
	}

	limit := 5
	res, err := a.appRepository.FindPublicApps(context.Background(), "", limit)
	a.Require().NoError(err)

	a.Require().Equal(limit, len(res))
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
	_, err := a.appInstallationRepository.Save(ctx, appChannel)
	a.Require().NoError(err)
}

func (a *AppDAOTestSuite) TestAppInstallationCreate() {
	appChannel := &appmodel.AppInstallation{
		ChannelID: channelID,
		AppID:     appID,
	}

	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, _ = a.appRepository.Save(ctx, app)
	_, err := a.appInstallationRepository.Create(ctx, appChannel)
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
	_, _ = a.appInstallationRepository.Save(ctx, appChannel)
	err := a.appInstallationRepository.DeleteByAppID(ctx, appID)
	a.Require().NoError(err)

	_, err = a.appInstallationRepository.Find(ctx, appChannel.ID())
	a.Require().True(apierr.IsNotFound(err))
}

func (a *AppDAOTestSuite) TestAppInstallationSaveAndFind() {
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

	_, err = a.appInstallationRepository.Save(ctx, appChannel)
	a.Require().NoError(err)

	res, err := a.appInstallationRepository.Find(ctx, installationID)

	a.Require().NoError(err)
	a.Require().NotNil(res)
}

func (a *AppDAOTestSuite) TestAppInstallationCreateAndFind() {
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

	_, err = a.appInstallationRepository.Create(ctx, appChannel)
	a.Require().NoError(err)

	res, err := a.appInstallationRepository.Find(ctx, installationID)

	a.Require().NoError(err)
	a.Require().NotNil(res)
	a.Require().Equal(installationID, res.ID())
}

func (a *AppDAOTestSuite) TestDuplicateInstallation() {
	appChannel := &appmodel.AppInstallation{
		ChannelID: channelID,
		AppID:     appID,
	}

	app := &appmodel.App{
		ID: appID,
	}

	ctx := context.Background()

	_, err := a.appRepository.Save(ctx, app)
	a.Require().NoError(err)

	_, err = a.appInstallationRepository.Create(ctx, appChannel)
	a.Require().NoError(err)

	_, err = a.appInstallationRepository.Create(ctx, appChannel)
	a.Require().True(apierr.IsConflict(err))
}

var testDisplay = &appmodel.AppDisplay{
	AppID:     appID,
	ManualURL: null.StringFrom("manual").Ptr(),
	DetailDescriptions: []map[string]any{{
		"test": `{ "test: test"}`,
	}},
	DetailImageURLs: []string{"image.url", "test.url"},
	I18nMap: map[string]appmodel.DisplayI18n{
		"ko": appmodel.DisplayI18n{
			DetailImageURLs: []string{"image.url", "ko.url"},
			DetailDescriptions: []map[string]any{{
				"test": `{ "test: ko"}`,
			}},
			ManualURL: null.StringFrom("ko-manual").Ptr(),
		},
	},
}

func (a *AppDAOTestSuite) TestSaveDisplay() {
	saved, err := a.displayRepo.Save(context.Background(), testDisplay)
	a.Require().NoError(err)
	a.Require().Equal(testDisplay.AppID, saved.AppID)
}

func (a *AppDAOTestSuite) TestFindDisplay() {
	_, err := a.displayRepo.Save(context.Background(), testDisplay)
	a.Require().NoError(err)

	display, err := a.displayRepo.Find(context.Background(), appID)
	a.Require().NoError(err)

	a.Require().Equal(testDisplay.AppID, display.AppID)
}

func (a *AppDAOTestSuite) TestDelete() {
	_, err := a.displayRepo.Save(context.Background(), testDisplay)
	a.Require().NoError(err)

	err = a.displayRepo.Delete(context.Background(), testDisplay.AppID)
	a.Require().NoError(err)

	_, err = a.displayRepo.Find(context.Background(), testDisplay.AppID)
	a.Require().True(apierr.IsNotFound(err))
}

func (a *AppDAOTestSuite) TestFindAll() {

	var appIDs []string
	cnt := 10
	for i := 0; i < cnt; i++ {
		newDisplay := *testDisplay
		newDisplay.AppID = "app-" + strconv.Itoa(i)

		saved, err := a.displayRepo.Save(context.Background(), &newDisplay)
		a.Require().NoError(err)

		appIDs = append(appIDs, saved.AppID)
	}

	res, err := a.displayRepo.FindAll(context.Background(), appIDs)
	a.Require().NoError(err)

	a.Require().Equal(len(appIDs), len(res))

}

func TestAppDAOs(t *testing.T) {
	suite.Run(t, new(AppDAOTestSuite))
}
