package svc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mocksvc "github.com/channel-io/ch-app-store/generated/mock/app/svc"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/lib/log"
)

const (
	appID = "abcdefg123"
)

type AppCrudSvcTestSuite struct {
	suite.Suite

	crudSvc             svc.AppLifecycleSvc
	querySvc            svc.AppQuerySvc
	appRepo             *mocksvc.AppRepository
	displayRepo         *mocksvc.AppDisplayRepository
	installQuerySvc     *svc.InstalledAppQuerySvc
	appInstallationRepo *mocksvc.AppInstallationRepository
	installSvc          svc.AppInstallSvc
}

func (a *AppCrudSvcTestSuite) SetupTest() {
	a.appRepo = mocksvc.NewAppRepository(a.T())
	a.displayRepo = mocksvc.NewAppDisplayRepository(a.T())

	a.appInstallationRepo = mocksvc.NewAppInstallationRepository(a.T())

	a.querySvc = svc.NewAppQuerySvcImpl(a.appRepo)
	a.installSvc = svc.NewAppInstallSvc(log.NewNoOpLogger(), a.appInstallationRepo, a.appRepo, nil, nil)
	a.installQuerySvc = svc.NewInstallQuerySvc(a.appInstallationRepo, a.appRepo, a.installSvc)
	a.crudSvc = svc.NewAppLifecycleSvcImpl(a.appRepo, a.displayRepo, a.installSvc, a.installQuerySvc, nil)
}

func (a *AppCrudSvcTestSuite) TestCreate() {
	app := &appmodel.App{
		Title: "test",
	}

	a.appRepo.EXPECT().Save(mock.Anything, app).Return(app, nil)

	ctx := context.Background()

	res, err := a.crudSvc.Create(ctx, app)

	// ID 비교는 Hex()를 사용하기 때문에 불가능
	assert.Equal(a.T(), app.Title, res.Title)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestUpdate() {
	before := &appmodel.App{
		ID:    appID,
		Title: "test",
	}

	update := &appmodel.App{
		ID:    appID,
		Title: "newTitle",
	}

	a.appRepo.EXPECT().Find(mock.Anything, appID).Return(before, nil)
	a.appRepo.EXPECT().Save(mock.Anything, update).Return(update, nil)

	ctx := context.Background()

	res, err := a.crudSvc.Update(ctx, update)

	assert.Equal(a.T(), update.Title, res.Title)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestDelete() {
	app := &appmodel.App{
		ID:    appID,
		Title: "test",
	}
	installations := []*appmodel.AppInstallation{
		{AppID: appID, ChannelID: "testChannel1"},
		{AppID: appID, ChannelID: "testChannel2"},
		{AppID: appID, ChannelID: "testChannel3"},
	}

	a.appRepo.EXPECT().Find(mock.Anything, appID).Return(app, nil)
	a.appInstallationRepo.EXPECT().FindAllByAppID(mock.Anything, appID).Return(installations, nil)
	for _, i := range installations {
		a.appInstallationRepo.EXPECT().Delete(mock.Anything, i.ID()).Return(nil)
	}

	a.appRepo.EXPECT().Delete(mock.Anything, appID).Return(nil)
	a.displayRepo.EXPECT().Delete(mock.Anything, appID).Return(nil)
	err := a.crudSvc.Delete(context.Background(), appID)

	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestRead() {
	app := &appmodel.App{
		ID:    appID,
		Title: "test",
	}

	a.appRepo.EXPECT().Find(mock.Anything, appID).Return(app, nil)

	ctx := context.Background()

	res, err := a.querySvc.Read(ctx, appID)

	assert.Equal(a.T(), appID, res.ID)
	assert.Nil(a.T(), err)
}

//func (a *AppCrudSvcTestSuite) TestReadPublicApps() {
//	apps := []*appmodel.App{
//		&appmodel.App{
//			ID:        "1",
//			Title:     "public app1",
//			IsPrivate: false,
//		},
//		&appmodel.App{
//			ID:        "2",
//			Title:     "public app2",
//			IsPrivate: false,
//		},
//		&appmodel.App{
//			ID:        "3",
//			Title:     "public app3",
//			IsPrivate: false,
//		},
//	}
//
//	a.appRepo.EXPECT().FindPublicApps(mock.Anything, mock.Anything, mock.Anything).Return(apps, nil)
//
//	ctx := context.Background()
//	res, err := a.querySvc.ReadPublicApps(ctx, "0", 500)
//
//	assert.Equal(a.T(), 3, len(res))
//	assert.False(a.T(), res[0].IsPrivate)
//	assert.False(a.T(), res[1].IsPrivate)
//	assert.False(a.T(), res[2].IsPrivate)
//	assert.Nil(a.T(), err)
//}

func (a *AppCrudSvcTestSuite) TestReadAllByAppIDs() {
	apps := []*appmodel.App{
		{
			ID:    "1",
			Title: "public app1",
		},
		{
			ID:    "3",
			Title: "public app3",
		},
		{
			ID:    "4",
			Title: "public app4",
		},
	}

	a.appRepo.EXPECT().FindAll(mock.Anything, []string{"1", "3", "4"}).Return(apps, nil)

	ctx := context.Background()
	res, err := a.querySvc.ReadAllByAppIDs(ctx, []string{"1", "3", "4"})

	assert.Equal(a.T(), 3, len(res))
	assert.Nil(a.T(), err)
}

func TestAppCrudSvcTestSuite(t *testing.T) {
	suite.Run(t, new(AppCrudSvcTestSuite))
}
