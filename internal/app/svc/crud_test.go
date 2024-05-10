package svc_test

import (
	"context"
	"testing"

	mocksvc "github.com/channel-io/ch-app-store/generated/mock/app/svc"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	appID = "abcdefg123"
)

type AppCrudSvcTestSuite struct {
	suite.Suite

	crudSvc             svc.AppLifecycleSvc
	querySvc            svc.AppQuerySvc
	appRepo             *mocksvc.AppRepository
	appInstallationRepo *mocksvc.AppInstallationRepository
}

func (a *AppCrudSvcTestSuite) SetupTest() {
	a.appRepo = mocksvc.NewAppRepository(a.T())
	a.appInstallationRepo = mocksvc.NewAppInstallationRepository(a.T())

	a.querySvc = svc.NewAppQuerySvcImpl(a.appRepo)
	a.crudSvc = svc.NewAppLifecycleSvc(a.appRepo, a.appInstallationRepo, nil)
}

func (a *AppCrudSvcTestSuite) TestCreate() {
	app := &appmodel.App{
		Title: "test",
	}

	a.appRepo.EXPECT().Save(mock.Anything, app).Return(app, nil)

	ctx := context.Background()

	res, err := a.crudSvc.Create(ctx, app)

	// ID 비교는 Hex()를 사용하기 때문에 불가능
	assert.Equal(a.T(), appmodel.AppStateEnabled, res.State)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestUpdate() {
	before := &appmodel.App{
		ID:        appID,
		Title:     "test",
		IsPrivate: false,
	}

	update := &appmodel.App{
		ID:        appID,
		Title:     "newTitle",
		IsPrivate: true,
	}

	a.appRepo.EXPECT().FindApp(mock.Anything, appID).Return(before, nil)
	a.appRepo.EXPECT().Save(mock.Anything, update).Return(update, nil)

	ctx := context.Background()

	res, err := a.crudSvc.Update(ctx, update)

	assert.True(a.T(), res.IsPrivate)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestDelete() {
	app := &appmodel.App{
		ID:        appID,
		Title:     "test",
		IsPrivate: false,
	}

	a.appRepo.EXPECT().FindApp(mock.Anything, appID).Return(app, nil)
	a.appInstallationRepo.EXPECT().DeleteByAppID(mock.Anything, appID).Return(nil)
	a.appRepo.EXPECT().Delete(mock.Anything, appID).Return(nil)

	ctx := context.Background()

	err := a.crudSvc.Delete(ctx, appID)

	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestRead() {
	app := &appmodel.App{
		ID:        appID,
		Title:     "test",
		IsPrivate: false,
	}

	a.appRepo.EXPECT().FindApp(mock.Anything, appID).Return(app, nil)

	ctx := context.Background()

	res, err := a.querySvc.Read(ctx, appID)

	assert.Equal(a.T(), appID, res.ID)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestReadPublicApps() {
	apps := []*appmodel.App{
		&appmodel.App{
			ID:        "1",
			Title:     "public app1",
			IsPrivate: false,
		},
		&appmodel.App{
			ID:        "2",
			Title:     "public app2",
			IsPrivate: false,
		},
		&appmodel.App{
			ID:        "3",
			Title:     "public app3",
			IsPrivate: false,
		},
	}

	a.appRepo.EXPECT().FindPublicApps(mock.Anything, mock.Anything, mock.Anything).Return(apps, nil)

	ctx := context.Background()
	res, err := a.querySvc.ReadPublicApps(ctx, "0", 500)

	assert.Equal(a.T(), 3, len(res))
	assert.False(a.T(), res[0].IsPrivate)
	assert.False(a.T(), res[1].IsPrivate)
	assert.False(a.T(), res[2].IsPrivate)
	assert.Nil(a.T(), err)
}

func (a *AppCrudSvcTestSuite) TestReadAllByAppIDs() {
	apps := []*appmodel.App{
		&appmodel.App{
			ID:        "1",
			Title:     "public app1",
			IsPrivate: false,
		},
		&appmodel.App{
			ID:        "3",
			Title:     "public app3",
			IsPrivate: false,
		},
		&appmodel.App{
			ID:        "4",
			Title:     "public app4",
			IsPrivate: true,
		},
	}

	a.appRepo.EXPECT().FindApps(mock.Anything, []string{"1", "3", "4"}).Return(apps, nil)

	ctx := context.Background()
	res, err := a.querySvc.ReadAllByAppIDs(ctx, []string{"1", "3", "4"})

	assert.Equal(a.T(), 3, len(res))
	assert.Nil(a.T(), err)
}

func TestAppCrudSvcTestSuite(t *testing.T) {
	suite.Run(t, new(AppCrudSvcTestSuite))
}
