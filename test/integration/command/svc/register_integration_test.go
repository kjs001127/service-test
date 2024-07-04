package svc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"

	commandmodel "github.com/channel-io/ch-app-store/internal/command/model"
	"github.com/channel-io/ch-app-store/internal/command/svc"
	. "github.com/channel-io/ch-app-store/test/integration"
)

const (
	appID = "badfase"
)

type CommandRegisterSvcTestSuite struct {
	suite.Suite

	testHelper *TestHelper

	registerSvc *svc.RegisterSvc
	commandRepo svc.CommandRepository
}

func (c *CommandRegisterSvcTestSuite) SetupTest() {
	c.testHelper = NewTestHelper(
		testOpts,
		fx.Populate(&c.registerSvc),
		fx.Populate(&c.commandRepo),
	)
	c.testHelper.TruncateAll()
}

func (c *CommandRegisterSvcTestSuite) TearDownSuite() {
	c.testHelper.Stop()
}

func (c *CommandRegisterSvcTestSuite) TestRegister() {
	ctx := context.Background()

	commands := []*commandmodel.Command{&commandmodel.Command{
		AppID:              appID,
		Name:               "test command",
		Scope:              "desk",
		ActionFunctionName: "gif",
	}}

	req := &svc.CommandRegisterRequest{
		AppID:    appID,
		Commands: commands,
	}

	err := c.registerSvc.Register(ctx, req)

	c.Require().NotNil(err)
}

func (c *CommandRegisterSvcTestSuite) TestDeregisterAll() {
	ctx := context.Background()

	commands := []*commandmodel.Command{&commandmodel.Command{
		AppID:              appID,
		Name:               "test command",
		Scope:              "desk",
		ActionFunctionName: "gif",
	}}

	req := &svc.CommandRegisterRequest{
		AppID:    appID,
		Commands: commands,
	}

	err := c.registerSvc.Register(ctx, req)

	c.Require().NotNil(err)

	err = c.registerSvc.DeregisterAll(ctx, appID)

	c.Require().Nil(err)
}

func TestCommandRegisterSvcSuite(t *testing.T) {
	suite.Run(t, new(CommandRegisterSvcTestSuite))
}
