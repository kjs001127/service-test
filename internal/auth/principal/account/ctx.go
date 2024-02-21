package account

import (
	"context"
	"errors"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	"github.com/channel-io/ch-app-store/internal/auth/principal"
)

type ContextAuthorizer interface {
	Authorize(
		ctx context.Context,
		channelContext app.ChannelContext,
		invoker ManagerPrincipal,
	) error
}

type ContextAuthorizerImpl struct {
	userFetcher   ManagerFetcher
	chatValidator principal.ChatValidator
}

func NewContextAuthorizer(userFetcher ManagerFetcher, chatValidator principal.ChatValidator) *ContextAuthorizerImpl {
	return &ContextAuthorizerImpl{userFetcher: userFetcher, chatValidator: chatValidator}
}

func (c ContextAuthorizerImpl) Authorize(
	ctx context.Context,
	channelContext app.ChannelContext,
	invoker ManagerPrincipal,
) error {

	if invoker.ChannelID != channelContext.Channel.ID {
		return apierr.Unauthorized(errors.New("channelID does not match"))
	}

	if err := c.chatValidator.ValidateChat(
		ctx,
		invoker.Token,
		principal.Chat{
			Type: channelContext.Chat.Type,
			ID:   channelContext.Chat.ID,
		},
	); err != nil {
		return err
	}

	return nil
}
