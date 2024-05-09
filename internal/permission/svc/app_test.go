package svc_test

import (
	"context"
	"testing"

	mocksvc "github.com/channel-io/ch-app-store/generated/mock/app/svc"
	mockrepo "github.com/channel-io/ch-app-store/generated/mock/permission/svc"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	account "github.com/channel-io/ch-app-store/internal/permission/model"
	"github.com/channel-io/ch-app-store/internal/permission/svc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	accountID = "123"
)

type AccountAppPermissionSvcTestSuite struct {
	suite.Suite

	permissionSvc  svc.AccountAppPermissionSvc
	appCrudSvc     *mocksvc.AppCrudSvc
	appAccountRepo *mockrepo.AppAccountRepo
}

func (s *AccountAppPermissionSvcTestSuite) SetupTest() {
	s.appCrudSvc = mocksvc.NewAppCrudSvc(s.T())
	s.appAccountRepo = mockrepo.NewAppAccountRepo(s.T())
	s.permissionSvc = svc.NewAccountAppPermissionSvc(s.appCrudSvc, s.appAccountRepo)
}

func (s *AccountAppPermissionSvcTestSuite) TestGetCallableApps() {
	appAccounts := []*account.AppAccount{
		&account.AppAccount{AppID: "1", AccountID: accountID},
		&account.AppAccount{AppID: "2", AccountID: accountID},
		&account.AppAccount{AppID: "3", AccountID: accountID},
	}

	privateApps := []*appmodel.App{
		&appmodel.App{ID: "1", Title: "App 1", IsPrivate: true},
		&appmodel.App{ID: "2", Title: "App 2", IsPrivate: true},
		&appmodel.App{ID: "3", Title: "App 3", IsPrivate: true},
	}

	publicApps := []*appmodel.App{
		&appmodel.App{ID: "4", Title: "App 4", IsPrivate: false},
		&appmodel.App{ID: "5", Title: "App 5", IsPrivate: false},
		&appmodel.App{ID: "6", Title: "App 6", IsPrivate: false},
	}

	s.appAccountRepo.EXPECT().FetchAllByAccountID(mock.Anything, accountID).Return(appAccounts, nil)

	s.appCrudSvc.EXPECT().ReadAllByAppIDs(mock.Anything, []string{"1", "2", "3"}).Return(privateApps, nil)
	s.appCrudSvc.EXPECT().ReadPublicApps(mock.Anything, mock.Anything, mock.Anything).Return(publicApps, nil)

	ctx := context.Background()
	apps, err := s.permissionSvc.GetCallableApps(ctx, accountID)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), apps)
}

func TestAccountAppPermissionSvcSuite(t *testing.T) {
	suite.Run(t, new(AccountAppPermissionSvcTestSuite))
}
