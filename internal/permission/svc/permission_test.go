package svc

import (
	"context"
	"testing"

	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	roleID = "843"
)

type PermissionUtilSuite struct {
	suite.Suite

	permissionUtil PermissionUtil
	roleFetcher    *mockaccount.ManagerRoleFetcher
}

func (p *PermissionUtilSuite) SetupTest() {
	p.roleFetcher = &mockaccount.ManagerRoleFetcher{}
	p.permissionUtil = NewPermissionUtil(p.roleFetcher)
}

func (p *PermissionUtilSuite) TestIsOwner() {
	role := account.ManagerRole{
		ID:       roleID,
		RoleType: "owner",
		Name:     "owner",
	}

	p.roleFetcher.EXPECT().FetchRole(mock.Anything, roleID).Return(role, nil)

	ctx := context.Background()
	manager := account.Manager{
		RoleID: roleID,
	}
	res := p.permissionUtil.isOwner(ctx, manager)

	assert.True(p.T(), res)
}

func TestPermissionUtilSuite(t *testing.T) {
	suite.Run(t, new(PermissionUtilSuite))
}
