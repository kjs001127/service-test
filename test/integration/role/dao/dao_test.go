package dao_test

import (
	"context"
	"testing"

	"github.com/channel-io/ch-app-store/internal/role/model"
	"github.com/channel-io/ch-app-store/internal/role/svc"
	. "github.com/channel-io/ch-app-store/test/integration"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

const (
	appID     = "1"
	channelID = "1"
	roleId    = "test"
	roleType  = model.RoleTypeChannel
	clientId  = "client"
	secret    = "secret"

	oldVersion = 1
	newVersion = 2
)

type AppDAOTestSuite struct {
	suite.Suite
	testApp *TestHelper

	agreementRepo svc.ChannelRoleAgreementRepo
	roleRepo      svc.AppRoleRepository
}

func (a *AppDAOTestSuite) SetupTest() {
	a.testApp = NewTestHelper(
		testOpts,
		fx.Populate(&a.agreementRepo),
		fx.Populate(&a.roleRepo),
	)
	a.testApp.TruncateAll()
}

func (a *AppDAOTestSuite) TearDownSuite() {
	a.testApp.Stop()
}

func (a *AppDAOTestSuite) TestRoleSave() {
	created, err := a.roleRepo.Save(context.Background(), &model.AppRole{
		AppID:   appID,
		RoleID:  roleId,
		Type:    roleType,
		Version: oldVersion,
		Credentials: &model.Credentials{
			ClientID:     clientId,
			ClientSecret: secret,
		},
	})

	a.Require().NoError(err)
	a.Require().NotNil(created)
	a.Require().NotEmpty(created.ID)
}

func (a *AppDAOTestSuite) TestLatestRoleFetch() {
	_, err := a.roleRepo.Save(context.Background(), &model.AppRole{
		AppID:   appID,
		RoleID:  roleId,
		Type:    roleType,
		Version: oldVersion,
		Credentials: &model.Credentials{
			ClientID:     clientId,
			ClientSecret: secret,
		},
	})
	a.Require().NoError(err)

	newer, err := a.roleRepo.Save(context.Background(), &model.AppRole{
		AppID:   appID,
		RoleID:  roleId,
		Type:    roleType,
		Version: newVersion,
		Credentials: &model.Credentials{
			ClientID:     clientId,
			ClientSecret: secret,
		},
	})

	a.Require().NoError(err)

	latest, err := a.roleRepo.FindLatestRole(context.Background(), appID, roleType)
	a.Require().NoError(err)

	a.Require().Equal(newer.ID, latest.ID)
	a.Require().Equal(newVersion, latest.Version)
}

func (a *AppDAOTestSuite) TestLatestRolesFetch() {

	for _, t := range model.AvailableRoleTypes {
		_, err := a.roleRepo.Save(context.Background(), &model.AppRole{
			AppID:   appID,
			RoleID:  roleId,
			Type:    t,
			Version: oldVersion,
			Credentials: &model.Credentials{
				ClientID:     clientId,
				ClientSecret: secret,
			},
		})
		a.Require().NoError(err)

		_, err = a.roleRepo.Save(context.Background(), &model.AppRole{
			AppID:   appID,
			RoleID:  roleId,
			Type:    t,
			Version: newVersion,
			Credentials: &model.Credentials{
				ClientID:     clientId,
				ClientSecret: secret,
			},
		})
		a.Require().NoError(err)
	}

	latestRoles, err := a.roleRepo.FindLatestRoles(context.Background(), []string{appID}, model.AvailableRoleTypes)
	a.Require().NoError(err)
	a.Require().Equal(len(model.AvailableRoleTypes), len(latestRoles))

	for _, latestRole := range latestRoles {
		a.Require().Equal(newVersion, latestRole.Version)
	}
}

func (a *AppDAOTestSuite) TestAgreementWithMultipleVersion() {
	for _, t := range model.AvailableRoleTypes {
		_, err := a.roleRepo.Save(context.Background(), &model.AppRole{
			AppID:   appID,
			RoleID:  roleId,
			Type:    t,
			Version: oldVersion,
			Credentials: &model.Credentials{
				ClientID:     clientId,
				ClientSecret: secret,
			},
		})
		a.Require().NoError(err)

		_, err = a.roleRepo.Save(context.Background(), &model.AppRole{
			AppID:   appID,
			RoleID:  roleId,
			Type:    t,
			Version: newVersion,
			Credentials: &model.Credentials{
				ClientID:     clientId,
				ClientSecret: secret,
			},
		})
		a.Require().NoError(err)
	}

	roles, err := a.agreementRepo.FindLatestUnAgreedRoles(context.Background(), channelID, []string{appID}, model.AvailableRoleTypes)
	a.Require().NoError(err)
	a.Require().Equal(len(model.AvailableRoleTypes), len(roles))

	for _, latestRole := range roles {
		a.Require().Equal(newVersion, latestRole.Version)
	}
}

func (a *AppDAOTestSuite) TestAgreementWithExistingAgreement() {
	appRole, err := a.roleRepo.Save(context.Background(), &model.AppRole{
		AppID:   appID,
		RoleID:  roleId,
		Type:    roleType,
		Version: 1,
		Credentials: &model.Credentials{
			ClientID:     clientId,
			ClientSecret: secret,
		},
	})
	a.Require().NoError(err)

	err = a.agreementRepo.Save(context.Background(), &model.ChannelRoleAgreement{
		ChannelID: channelID,
		AppRoleID: appRole.ID,
	})
	a.Require().NoError(err)

	roles, err := a.agreementRepo.FindLatestUnAgreedRoles(
		context.Background(),
		channelID,
		[]string{appID},
		[]model.RoleType{roleType},
	)
	a.Require().NoError(err)

	a.Require().Equal(0, len(roles))
}

func TestAppDAOs(t *testing.T) {
	suite.Run(t, new(AppDAOTestSuite))
}
