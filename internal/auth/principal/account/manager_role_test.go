package account_test

import (
	"context"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/test"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

type ManagerRoleFetcherSuite struct {
	suite.Suite

	testApp *test.TestApp

	fetcher account.ManagerRoleFetcher
}

func (m *ManagerRoleFetcherSuite) SetupSuite() {
	m.testApp = test.NewTestApp(
		fx.Populate(&m.fetcher),
	)
}

func (m *ManagerRoleFetcherSuite) TestFetchRole() {

	res, _ := m.fetcher.FetchRole(context.Background(), "1", "843")

	assert.Equal(m.T(), res.RoleType, "owner")
}
