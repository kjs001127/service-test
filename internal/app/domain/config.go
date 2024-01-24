package domain

import (
	"context"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
)

type ConfigMap map[string]string
type ConfigSchemas []ConfigSchema

type ConfigSchema struct {
	Name       string
	Type       string
	Key        string
	Default    null.String
	Help       null.String
	Attributes map[string]any
}

func (s ConfigSchemas) Default() ConfigMap {
	ret := make(map[string]string)
	for _, schema := range s {
		if schema.Default.Valid {
			ret[schema.Key] = schema.Default.String
		}
	}
	return ret
}

type ConfigSvc struct {
	repo    AppRepository
	checker HttpAppChecker
}

func NewConfigSvc(repo AppRepository) *ConfigSvc {
	return &ConfigSvc{repo: repo}
}

func (s *ConfigSvc) CheckConfig(ctx context.Context, install InstallInfo, input ConfigMap) (CheckReturn, error) {
	app, err := s.repo.Fetch(ctx, install.AppID)
	if err != nil {
		return CheckReturn{}, errors.Wrap(err, "app fetch fail")
	}

	if !app.CheckURL.Valid {
		return CheckReturn{}, apierr.BadRequest(errors.New("app checkUrl is empty"))
	}

	return s.checker.Request(
		ctx,
		app.CheckURL.String,
		CheckRequest{
			Type: CheckTypeConfig,
			Data: input,
		},
	)
}

func (s *ConfigSvc) DefaultConfigOf(ctx context.Context, install InstallInfo) (ConfigMap, error) {
	app, err := s.repo.Fetch(ctx, install.AppID)
	if err != nil {
		return nil, errors.Wrap(err, "app fetch fail")
	}

	return app.ConfigSchemas.Default(), nil
}

type CheckType string

const (
	CheckTypeConfig = CheckType("config")
)

type CheckRequest struct {
	ChannelId string
	Type      CheckType
	Data      any
}

type CheckReturn struct {
	Success  bool
	Messages map[string]any
}

type HttpAppChecker interface {
	Request(ctx context.Context, url string, req CheckRequest) (CheckReturn, error)
}
