package principal

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ChatValidator interface {
	ValidateChat(ctx context.Context, token Token, chat Chat) error
}

type ChatValidatorImpl struct {
	cli         *resty.Client
	chatAuthUrl string
}

func NewChatValidator(cli *resty.Client, chatAuthUrl string) *ChatValidatorImpl {
	return &ChatValidatorImpl{cli: cli, chatAuthUrl: chatAuthUrl}
}

type Chat struct {
	Type string
	ID   string
}

func (c *ChatValidatorImpl) ValidateChat(ctx context.Context, token Token, chat Chat) error {
	req := c.cli.R()
	req.SetContext(ctx)
	req.Header.Set(token.Type().Header(), token.Value())

	resp, err := req.Get(c.chatAuthUrl + fmt.Sprintf("/%ss/%s", chat.Type, chat.ID))
	if err != nil {
		return err
	}

	if !isSuccess(resp.StatusCode()) {
		return errors.New("chat auth failed")
	}

	return nil
}

func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
