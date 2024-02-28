package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/native/domain"
)

const (
	version           = "v1"
	v1BaseUri         = "/api/admin/core/" + version
	messageBaseUri    = v1BaseUri + "/messages"
	managerBaseUri    = v1BaseUri + "/managers"
	userChatBaseUri   = v1BaseUri + "/user-chats"
	contentTypeHeader = "Content-Type"
	mimeTypeJson      = "application/json"
)

type CoreApi struct {
	adminUrl  string
	resty     *resty.Client
	urlRouter map[string]string
}

func NewCoreApi(adminUrl string, resty *resty.Client) *CoreApi {
	api := &CoreApi{adminUrl: adminUrl, resty: resty}
	api.urlRouter = map[string]string{
		"writeUserChatMessage": messageBaseUri + "/writeUserChatMessage",
		"writeGroupMessage":    messageBaseUri + "/writeGroupMessage",

		"getManager":       managerBaseUri + "/getManager",
		"batchGetManagers": managerBaseUri + "/batchGetManagers",
		"searchManagers":   managerBaseUri + "/searchManagers",

		"getUserChat": userChatBaseUri + "/getUserChat",
	}
	return api
}

func (a *CoreApi) ListMethods() []string {
	var methods []string
	for method, _ := range a.urlRouter {
		methods = append(methods, method)
	}
	return methods
}

func (a *CoreApi) Handle(ctx context.Context, fnReq domain.NativeFunctionRequest) domain.NativeFunctionResponse {

	uri, ok := a.urlRouter[fnReq.Method]
	if !ok {
		return domain.WrapCommonErr(errors.New("mapping not found on core api handler"))
	}

	marshaled, err := json.Marshal(fnReq.Params)
	if err != nil {
		return domain.WrapCommonErr(err)
	}

	r := a.resty.R()

	if len(fnReq.Token.Type) > 0 && len(fnReq.Token.Value) > 0 {
		r.SetHeader(fnReq.Token.Type, fnReq.Token.Value)
	}

	r.SetHeader(contentTypeHeader, mimeTypeJson)
	r.SetBody(marshaled)
	r.SetContext(ctx)

	resp, err := r.Post(a.adminUrl + uri)
	if err != nil {
		return domain.WrapCommonErr(err)
	}

	if resp.IsError() {
		return domain.WrapCommonErr(fmt.Errorf("request failed, body: %s", resp.Body()))
	}

	return domain.NativeFunctionResponse{
		Result: resp.Body(),
	}
}
