package domain

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const bufSize = 1024 * 2

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, bufSize)
	},
}

type RemoteApp struct {
	app.AppData

	RoleID   string
	ClientID string
	Secret   string

	HookURL     null.String
	FunctionURL null.String
	WamURL      null.String
	CheckURL    null.String

	requester HttpRequester
}

const jsonRpcVersion = "2.0"

type JsonRpcRequest struct {
	JsonRpc string
	ID      string
	Method  string
	Params  json.RawMessage
}

type JsonRpcResponse struct {
	JsonRpc string
	ID      string
	Result  json.RawMessage
}

type HttpRequest struct {
	Method      string
	Url         string
	ContentType string
	Body        []byte
	Headers     map[string]string
}

type HttpRequester interface {
	Request(ctx context.Context, request HttpRequest) (io.ReadCloser, error)
}

func (a *RemoteApp) Invoke(ctx context.Context, function string, input null.JSON) (null.JSON, error) {
	id := uid.New().Hex()

	jsonReq := JsonRpcRequest{
		JsonRpc: jsonRpcVersion,
		ID:      id,
		Method:  function,
		Params:  json.RawMessage(input.JSON),
	}

	marshaled, err := json.Marshal(jsonReq)
	if err != nil {
		return null.JSON{}, err
	}

	if !a.FunctionURL.Valid {
		return null.JSON{}, apierr.BadRequest(errors.New("function url invalid"))
	}

	reader, err := a.requester.Request(ctx, HttpRequest{
		Body:   marshaled,
		Method: http.MethodPost,
		Url:    a.FunctionURL.String,
	})

	ret, err := io.ReadAll(reader)
	if err != nil {
		return null.JSON{}, err
	}

	if err := reader.Close(); err != nil {
		return null.JSON{}, err
	}

	var resp JsonRpcResponse
	if err := json.Unmarshal(ret, &resp); err != nil {
		return null.JSON{}, err
	}

	return null.JSONFrom(resp.Result), nil
}

func (a *RemoteApp) Data() *app.AppData {
	return &a.AppData
}

func (a *RemoteApp) CheckInstallable(ctx context.Context, channelID string) error {
	return nil
}

func (a *RemoteApp) OnInstall(ctx context.Context, channelID string) error {
	return nil
}

func (a *RemoteApp) OnUnInstall(ctx context.Context, channelID string) error {
	return nil
}

func (a *RemoteApp) OnConfigSet(ctx context.Context, channelID string, input app.ConfigMap) error {
	return nil
}

func (a *RemoteApp) StreamFile(ctx context.Context, path string, writer io.Writer) error {
	if !a.WamURL.Valid {
		return apierr.BadRequest(errors.New("wam url invalid"))
	}

	url := a.WamURL.String + path

	reader, err := a.requester.Request(ctx, HttpRequest{
		Method: http.MethodGet,
		Url:    url,
	})
	if err != nil {
		return err
	}
	defer reader.Close()

	return doStream(reader, writer)
}

func doStream(from io.ReadCloser, to io.Writer) error {
	buf := bufPool.Get().([]byte)
	defer bufPool.Put(buf)

	var n int
	var err error
	for ; err == nil; n, err = from.Read(buf) {
		if n <= 0 {
			continue
		}

		if _, err := to.Write(buf[:n]); err != nil {
			return err
		}
	}

	if err != io.EOF {
		return err
	}

	return nil
}

type RemoteAppRepository interface {
	Index(ctx context.Context, since string, limit int) ([]*RemoteApp, error)
	Fetch(ctx context.Context, appID string) (*RemoteApp, error)
	FindAll(ctx context.Context, appIDs []string) ([]*RemoteApp, error)
	Save(ctx context.Context, app *RemoteApp) (*RemoteApp, error)
	Update(ctx context.Context, app *RemoteApp) (*RemoteApp, error)
	Delete(ctx context.Context, appID string) error
}

type AppRepositoryAdapter struct {
	appRepository RemoteAppRepository
	requester     HttpRequester
}

func NewAppRepositoryAdapter(appRepository RemoteAppRepository, requester HttpRequester) *AppRepositoryAdapter {
	return &AppRepositoryAdapter{appRepository: appRepository, requester: requester}
}

func (i *AppRepositoryAdapter) Index(ctx context.Context, since string, limit int) ([]app.App, error) {
	apps, err := i.appRepository.Index(ctx, since, limit)
	if err != nil {
		return nil, err
	}

	ret := make([]app.App, len(apps))
	for _, a := range apps {
		a.requester = i.requester
		ret = append(ret, a)
	}

	return ret, nil
}

func (i *AppRepositoryAdapter) FindApps(ctx context.Context, appIDs []string) ([]app.App, error) {
	apps, err := i.appRepository.FindAll(ctx, appIDs)
	if err != nil {
		return nil, err
	}

	ret := make([]app.App, len(apps))
	for _, a := range apps {
		a.requester = i.requester
		ret = append(ret, a)
	}

	return ret, nil
}

func (i *AppRepositoryAdapter) FindApp(ctx context.Context, appID string) (app.App, error) {
	one, err := i.appRepository.Fetch(ctx, appID)
	if err != nil {
		return nil, err
	}
	one.requester = i.requester
	return one, nil
}
