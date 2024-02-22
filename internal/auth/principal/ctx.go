package principal

import (
	"context"

	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

type CommandCtxAuthorizer interface {
	Authorize(
		ctx context.Context,
		channelContext cmd.CommandContext,
		token Token,
	) error
}

type CommandCtxAuthorizerImpl struct {
	chatValidator ChatValidator
}

func NewCommandCtxAuthorizer(chatValidator ChatValidator) *CommandCtxAuthorizerImpl {
	return &CommandCtxAuthorizerImpl{chatValidator: chatValidator}
}

func (c CommandCtxAuthorizerImpl) Authorize(
	ctx context.Context,
	cmdCtx cmd.CommandContext,
	token Token,
) error {
	if err := c.chatValidator.ValidateChat(
		ctx,
		token,
		Chat{
			Type: cmdCtx.Chat.Type,
			ID:   cmdCtx.Chat.ID,
		},
	); err != nil {
		return err
	}
	return nil
}
