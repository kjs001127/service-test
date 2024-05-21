package integration_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"

	mockappsvc "github.com/channel-io/ch-app-store/generated/mock/app/svc"
	mockcmdsvc "github.com/channel-io/ch-app-store/generated/mock/command/svc"
	"github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	cmdsvc "github.com/channel-io/ch-app-store/internal/command/svc"
	hookmodel "github.com/channel-io/ch-app-store/internal/hook/model"
	hooksvc "github.com/channel-io/ch-app-store/internal/hook/svc"
	managersvc "github.com/channel-io/ch-app-store/internal/manager/svc"
	. "github.com/channel-io/ch-app-store/test/integration"
)

const (
	appID     = "1"
	accountID = "123"
	channelID = "1"
	lang      = "ko"
)

var (
	toggleFnName    = "toggle"
	unInstallFnName = "uninstall"
	installFnName   = "install"
)

var testApp = &model.App{
	ID:    appID,
	Title: "test app",
}

var testManager = account.Manager{
	AccountID: accountID,
	ChannelID: channelID,
	Language:  lang,
}

var installID = model.InstallationID{
	AppID:     appID,
	ChannelID: channelID,
}

type HookTestSuite struct {
	suite.Suite
	helper *TestHelper

	toggleInvoker  mockappsvc.TypedInvoker[hooksvc.ToggleHookRequest, hooksvc.ToggleHookResponse]
	installInvoker mockappsvc.Invoker
	installSvc     mockappsvc.AppInstallSvc
	toggleSvc      mockcmdsvc.ToggleSvc

	managerToggleSvc  *managersvc.ManagerAwareToggleSvc
	managerInstallSvc *managersvc.ManagerAwareInstallSvc

	hookRepo    hooksvc.ToggleHookRepository
	installRepo hooksvc.InstallHookRepository
}

func (p *HookTestSuite) SetupTest() {
	p.helper = NewTestHelper(
		testOpts,
		Mock[svc.AppInstallSvc](&p.installSvc),
		Mock[cmdsvc.ToggleSvc](&p.toggleSvc),
		Mock[svc.TypedInvoker[hooksvc.ToggleHookRequest, hooksvc.ToggleHookResponse]](&p.toggleInvoker),
		Mock[svc.Invoker](&p.installInvoker),

		fx.Populate(&p.managerInstallSvc),
		fx.Populate(&p.managerToggleSvc),
		fx.Populate(&p.hookRepo),
		fx.Populate(&p.installRepo),
	)
	p.helper.TruncateAll()
}

func (p *HookTestSuite) TearDownSuite() {
	p.helper.Stop()
}

func (p *HookTestSuite) TestToggleHookCalled() {
	_ = p.hookRepo.Save(context.Background(), &hookmodel.CommandToggleHooks{
		AppID:              appID,
		ToggleFunctionName: toggleFnName,
	})

	p.toggleSvc.EXPECT().
		Toggle(mock.Anything, installID, true).
		Return(nil).
		Once()

	var captured svc.TypedRequest[hooksvc.ToggleHookRequest]
	p.toggleInvoker.EXPECT().
		Invoke(mock.Anything, mock.Anything, mock.Anything).
		Run(func(ctx context.Context, appID string, req svc.TypedRequest[hooksvc.ToggleHookRequest]) {
			captured = req
		}).
		Return(svc.TypedResponse[hooksvc.ToggleHookResponse]{
			Result: hooksvc.ToggleHookResponse{
				Enable: true,
			},
		}).Once()

	err := p.managerToggleSvc.Toggle(context.Background(), testManager, installID, true)

	fmt.Println(captured.Params.AppID)
	p.Require().NoError(err)
	p.Require().Equal(toggleFnName, captured.FunctionName)
	p.Require().Equal(appID, captured.Params.AppID)
	p.Require().Equal(channelID, captured.Params.ChannelID)
	p.Require().Equal(testManager.Language, captured.Params.Language)
}

func (p *HookTestSuite) TestToggleHookRejected() {
	_ = p.hookRepo.Save(context.Background(), &hookmodel.CommandToggleHooks{
		AppID:              appID,
		ToggleFunctionName: toggleFnName,
	})

	p.toggleSvc.EXPECT().
		Toggle(mock.Anything, installID, true).
		Return(nil).
		Once()

	p.toggleInvoker.EXPECT().
		Invoke(mock.Anything, mock.Anything, mock.Anything).
		Return(svc.TypedResponse[hooksvc.ToggleHookResponse]{
			Error: &svc.Error{
				Type:    "test",
				Message: "testErrMsg",
			},
		}).Once()

	err := p.managerToggleSvc.Toggle(context.Background(), testManager, installID, true)
	p.Require().Error(err)
}

func (p *HookTestSuite) TestToggleHookNotExists() {
	p.toggleSvc.EXPECT().
		Toggle(mock.Anything, installID, true).
		Return(nil).
		Once()

	err := p.managerToggleSvc.Toggle(context.Background(), testManager, installID, true)

	p.Require().NoError(err)
}

func (p *HookTestSuite) TestInstallHookNotExists() {
	p.installSvc.EXPECT().
		InstallAppById(mock.Anything, installID).
		Return(testApp, nil)

	ret, err := p.managerInstallSvc.Install(context.Background(), testManager, installID)

	p.Require().NotNil(ret)
	p.Require().NoError(err)
}

func (p *HookTestSuite) TestInstallHookCalled() {
	_ = p.installRepo.Save(context.Background(), appID, &hookmodel.AppInstallHooks{
		AppID:               appID,
		InstallFunctionName: &installFnName,
	})

	p.installSvc.EXPECT().
		InstallAppById(mock.Anything, installID).
		Return(testApp, nil)

	wait := make(chan any)
	p.installInvoker.EXPECT().
		Invoke(mock.Anything, appID, mock.Anything).
		Run(func(ctx context.Context, appID string, req svc.JsonFunctionRequest) {
			close(wait)
		}).
		Return(svc.JsonFunctionResponse{Result: []byte("{}")}).
		Once()

	ret, err := p.managerInstallSvc.Install(context.Background(), testManager, installID)
	<-wait

	p.Require().NotNil(ret)
	p.Require().NoError(err)
}

func (p *HookTestSuite) TestUninstallHookCalled() {
	_ = p.installRepo.Save(context.Background(), appID, &hookmodel.AppInstallHooks{
		AppID:                 appID,
		UninstallFunctionName: &unInstallFnName,
	})

	p.installSvc.EXPECT().
		UnInstallApp(mock.Anything, installID).
		Return(nil)

	p.installInvoker.EXPECT().
		Invoke(mock.Anything, appID, mock.Anything).
		Once()

	err := p.managerInstallSvc.UnInstall(context.Background(), testManager, installID)

	p.Require().NoError(err)
}

func TestHooks(t *testing.T) {
	suite.Run(t, new(HookTestSuite))
}
