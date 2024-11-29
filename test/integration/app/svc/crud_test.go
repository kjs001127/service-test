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

type CrudTestSuite struct {
	suite.Suite

	testHelper *TestHelper

	appLifecycleSvc svc.AppLifecycleSvc
	appQuerySvc     svc.AppQuerySvc
}

func (a *CrudTestSuite) SetupTest() {
	a.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&a.appLifecycleSvc),
		fx.Populate(&a.appQuerySvc),
	)
	a.testHelper.TruncateAll()
}

func (a *CrudTestSuite) TearDownSuite() {
	a.testHelper.Stop()
}

func (a *CrudTestSuite) TestAppCreate() {
	app, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})

	a.Require().NoError(err)
	a.Require().NotNil(app)
	a.Require().NotEmpty(app.ID)
	a.Require().Equal(app.Title, "test app")
}

func (a *CrudTestSuite) TestAppRead() {
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

func (a *CrudTestSuite) TestAppUpdate() {
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

func (a *CrudTestSuite) TestAppDelete() {
	created, err := a.appLifecycleSvc.Create(context.Background(), &appmodel.App{
		Title: "test app",
	})
	if err != nil {
		panic(err)
	}
	err = a.appLifecycleSvc.Delete(context.Background(), created.ID)

	a.Require().NoError(err)
}

func (a *CrudTestSuite) TestReadAllByAppIDs() {
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

func TestAppIntegrationSuite(t *testing.T) {
	suite.Run(t, new(CrudTestSuite))
}
