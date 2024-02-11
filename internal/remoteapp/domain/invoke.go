package domain

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/uid"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

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

type Invoker struct {
	requester HttpRequester
	repo      AppUrlRepository
}

func NewInvoker(requester HttpRequester, repo AppUrlRepository) *Invoker {
	return &Invoker{requester: requester, repo: repo}
}

func (a *Invoker) Invoke(ctx context.Context, request app.FunctionRequest, out app.FunctionResponse) error {
	urls, err := a.repo.Fetch(ctx, request.AppID)
	if err != nil {
		return err
	}

	if urls.FunctionURL == nil {
		return apierr.BadRequest(errors.New("function url invalid"))
	}

	jsonRPCReq, err := a.toJsonRPCRequest(request)
	if err != nil {
		return err
	}

	reader, err := a.requester.Request(ctx, HttpRequest{
		Body:   jsonRPCReq,
		Method: http.MethodPost,
		Url:    *urls.FunctionURL,
	})
	if err != nil {
		return err
	}

	defer reader.Close()

	ret, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	if err := a.fromJsonRPCResponse(ret, out); err != nil {
		return err
	}
	return nil
}

func (a *Invoker) toJsonRPCRequest(request app.FunctionRequest) ([]byte, error) {
	id := uid.New().Hex()

	jsonReq := make(map[string]any)
	jsonReq["id"] = id
	jsonReq["jsonrpc"] = jsonRpcVersion
	jsonReq["method"] = request.FunctionName
	jsonReq["params"] = request.Body
	jsonReq["context"] = request.Context
	jsonReq["caller"] = request.Caller

	return json.Marshal(jsonReq)
}

func (a *Invoker) fromJsonRPCResponse(ret []byte, out app.FunctionResponse) error {
	var jsonResp JsonRpcResponse
	if err := json.Unmarshal(ret, &jsonResp); err != nil {
		return err
	}

	if err := json.Unmarshal(jsonResp.Result, &out); err != nil {
		return err
	}
	return nil
}
