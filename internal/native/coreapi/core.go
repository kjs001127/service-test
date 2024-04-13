package coreapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/internal/native"
)

const (
	version         = "v1"
	v1BaseUri       = "/api/admin/core/" + version
	messageBaseUri  = v1BaseUri + "/messages"
	managerBaseUri  = v1BaseUri + "/managers"
	userChatBaseUri = v1BaseUri + "/user-chats"

	contentTypeHeader = "Content-Type"
	mimeTypeJson      = "application/json"
)

type CoreApi struct {
	adminUrl  string
	resty     *resty.Client
	urlRouter map[string]string
}

func (a *CoreApi) RegisterTo(registry native.FunctionRegistry) {
	for method, _ := range a.urlRouter {
		registry.Register(method, a.Handle)
	}
}

func NewCoreApi(adminUrl string, resty *resty.Client) *CoreApi {
	api := &CoreApi{adminUrl: adminUrl, resty: resty}
	api.urlRouter = map[string]string{
		"writeGroupMessage":          messageBaseUri + "/writeGroupMessage",
		"writeGroupMessageAsManager": messageBaseUri + "/writeGroupMessageAsManager",

		"writeUserChatMessage":          messageBaseUri + "/writeUserChatMessage",
		"writeUserChatMessageAsManager": messageBaseUri + "/writeUserChatMessageAsManager",
		"writeUserChatMessageAsUser":    messageBaseUri + "/writeUserChatMessageAsUser",

		"writeDirectChatMessageAsManager": messageBaseUri + "/writeDirectChatMessageAsManager",
		"createDirectChat":                messageBaseUri + "/createDirectChat",

		"getManager":       managerBaseUri + "/getManager",
		"batchGetManagers": managerBaseUri + "/batchGetManagers",
		"searchManagers":   managerBaseUri + "/searchManagers",

		"getUserChat": userChatBaseUri + "/getUserChat",
	}
	return api
}

func (a *CoreApi) Handle(ctx context.Context, token native.Token, fnReq native.FunctionRequest) native.FunctionResponse {

	uri, ok := a.urlRouter[fnReq.Method]
	if !ok {
		return native.WrapCommonErr(errors.New("mapping not found on core api handler"))
	}

	r := a.resty.R()

	if len(token.Type) > 0 && len(token.Value) > 0 {
		r.SetHeader(token.Type, token.Value)
	}

	r.SetHeader(contentTypeHeader, mimeTypeJson)
	r.SetBody(fnReq.Params)
	r.SetContext(ctx)

	resp, err := r.Post(a.adminUrl + uri)
	if err != nil {
		return native.WrapCommonErr(err)
	}

	if resp.IsError() {
		return native.WrapCommonErr(fmt.Errorf("request failed, body: %s", resp.Body()))
	}

	return native.FunctionResponse{
		Result: resp.Body(),
	}
}