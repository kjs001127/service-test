package principal

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const uri = "/api/admin/auth"

type ChatValidator interface {
	ValidateChat(ctx context.Context, token Token, chat Chat) error
}

type ChatValidatorImpl struct {
	cli     *resty.Client
	baseUrl string
}

func NewChatValidator(cli *resty.Client, baseUrl string) *ChatValidatorImpl {
	return &ChatValidatorImpl{cli: cli, baseUrl: baseUrl}
}

type Chat struct {
	Type string
	ID   string
}

func (c *ChatValidatorImpl) ValidateChat(ctx context.Context, token Token, chat Chat) error {
	req := c.cli.R()
	req.SetContext(ctx)
	req.Header.Set(token.Type(), token.Value())

	resp, err := req.Get(c.baseUrl + uri + fmt.Sprintf("/%ss/%s", chat.Type, chat.ID))
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
