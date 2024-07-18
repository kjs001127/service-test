package permission_test

import (
	"context"
	"github.com/channel-io/ch-app-store/internal/permission/svc"
	. "github.com/channel-io/ch-app-store/test/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
	"testing"
)

type AppAccountRepoTestSuite struct {
	suite.Suite
	helper *TestHelper

	appAccountRepo svc.AppAccountRepo
}

func (a *AppAccountRepoTestSuite) SetupTest() {
	a.helper = NewTestHelper(
		testOpts,
		fx.Populate(&a.appAccountRepo),
	)

	a.helper.TruncateAll()
}

func (a *AppAccountRepoTestSuite) TearDownSuite() {
	a.helper.Stop()
}

func (a *AppAccountRepoTestSuite) TestDeleteByAppID() {
	appID := "test"
	accountID := "test"

	ctx := context.Background()
	a.appAccountRepo.Save(ctx, appID, accountID)

	err := a.appAccountRepo.DeleteByAppID(ctx, appID)

	assert.Nil(a.T(), err)
}

func TestAppAccountSuite(t *testing.T) {
	suite.Run(t, new(AppAccountRepoTestSuite))
}
