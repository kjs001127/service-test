package domain_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/mock"

	mockdomain "github.com/channel-io/ch-app-store/generated/mock/command/domain"
	"github.com/channel-io/ch-app-store/internal/command/domain"
	"github.com/channel-io/ch-app-store/lib/db/tx"
)

func init() {
	tx.SetDB(&sql.DB{})
}

func mockRepository(t *testing.T) *mockdomain.CommandRepository {
	appCmds := map[string][]*domain.Command{
		"1": {
			{
				ID:    "1",
				AppID: "1",
				Name:  "frontCommand",
				Scope: domain.ScopeFront,
			},
			{
				ID:    "2",
				AppID: "1",
				Name:  "deskCommand",
				Scope: domain.ScopeDesk,
			},
		},
		"2": {
			{
				ID:    "3",
				AppID: "2",
				Name:  "deskCommand2",
				Scope: domain.ScopeDesk,
			},
			{
				ID:    "4",
				AppID: "2",
				Name:  "frontCommand2",
				Scope: domain.ScopeFront,
			},
		},
	}

	repo := mockdomain.NewCommandRepository(t)
	for appID, cmds := range appCmds {
		repo.EXPECT().
			FetchAllByAppIDs(mock.Anything, appID).
			Return(cmds, nil).
			Maybe()
	}
	return repo
}

func TestNewInsert(t *testing.T) {
	mockRepo := mockRepository(t)
	svc := domain.NewRegisterService(mockRepo, domain.NewParamValidator())
	insertCmds := []*domain.Command{
		{
			AppID: "3",
			Name:  "deskCommand",
			Scope: domain.ScopeDesk,
		},
		{
			AppID: "3",
			Name:  "frontCommand",
			Scope: domain.ScopeFront,
		},
	}
	for _, cmd := range insertCmds {
		mockRepo.EXPECT().Save(mock.Anything, cmd).Once()
	}

	if err := svc.Register(context.Background(), "3", insertCmds); err != nil {
		t.Error(err)
	}

}
