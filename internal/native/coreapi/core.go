package coreapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/channel-io/ch-app-store/internal/native"
	fpa "github.com/channel-io/ch-app-store/internal/native/coreapi/action/firstparty"
	tpa "github.com/channel-io/ch-app-store/internal/native/coreapi/action/thirdparty"

	"github.com/go-resty/resty/v2"
)

const (
	version           = "v1"
	v1BaseUri         = "/api/admin/core/" + version
	messageBaseUri    = v1BaseUri + "/messages"
	managerBaseUri    = v1BaseUri + "/managers"
	userChatBaseUri   = v1BaseUri + "/user-chats"
	directChatBaseUri = v1BaseUri + "/direct-chats"
	userBaseUri       = v1BaseUri + "/users"
	channelBaseUri    = v1BaseUri + "/channels"
	mediumBaseUri     = v1BaseUri + "/mediums"
	groupBaseUri      = v1BaseUri + "/groups"

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
		tpa.WriteGroupMessage:          messageBaseUri + "/" + tpa.WriteGroupMessage,
		tpa.WriteGroupMessageAsManager: messageBaseUri + "/" + tpa.WriteGroupMessageAsManager,

		tpa.WriteUserChatMessage:            messageBaseUri + "/" + tpa.WriteUserChatMessage,
		tpa.WriteUserChatMessageAsManager:   messageBaseUri + "/" + tpa.WriteUserChatMessageAsManager,
		tpa.WriteUserChatMessageAsUser:      messageBaseUri + "/" + tpa.WriteUserChatMessageAsUser,
		tpa.WriteDirectChatMessageAsManager: messageBaseUri + "/" + tpa.WriteDirectChatMessageAsManager,

		fpa.CreateDirectChat: directChatBaseUri + "/" + fpa.CreateDirectChat,

		tpa.ManageUserChat: userChatBaseUri + "/" + tpa.ManageUserChat,

		tpa.GetManager:       managerBaseUri + "/" + tpa.GetManager,
		tpa.BatchGetManagers: managerBaseUri + "/" + tpa.BatchGetManagers,
		tpa.SearchManagers:   managerBaseUri + "/" + tpa.SearchManagers,

		tpa.GetUserChat: userChatBaseUri + "/" + tpa.GetUserChat,

		tpa.GetUser:    userBaseUri + "/" + tpa.GetUser,
		tpa.GetChannel: channelBaseUri + "/" + tpa.GetChannel,
		tpa.GetGroup:   groupBaseUri + "/" + tpa.GetGroup,
		fpa.SearchUser: userBaseUri + "/" + fpa.SearchUser,

		fpa.FindOrCreateContactAndUser: mediumBaseUri + "/" + fpa.FindOrCreateContactAndUser,
		fpa.SearchUserChatsByContact:   userChatBaseUri + "/" + fpa.SearchUserChatsByContact,
		fpa.UpdateUserChatState:        userChatBaseUri + "/" + fpa.UpdateUserChatState,
	}
	return api
}

func (a *CoreApi) Handle(ctx context.Context, token native.Token, fnReq native.FunctionRequest) native.FunctionResponse {

	uri, ok := a.urlRouter[fnReq.Method]
	if !ok {
		return native.WrapCommonErr(errors.New("mapping not found on core api handler"))
	}

	r := a.resty.R()

	if token.Exists && len(token.Value) > 0 {
		r.SetHeader("x-access-token", token.Value)
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
