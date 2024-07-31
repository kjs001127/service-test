package coreapi

import (
	"context"
	"errors"
	"fmt"

	"github.com/channel-io/ch-app-store/internal/native"
	pvt "github.com/channel-io/ch-app-store/internal/native/coreapi/action/private"
	pub "github.com/channel-io/ch-app-store/internal/native/coreapi/action/public"

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
		pub.WriteGroupMessage:          messageBaseUri + "/" + pub.WriteGroupMessage,
		pub.WriteGroupMessageAsManager: messageBaseUri + "/" + pub.WriteGroupMessageAsManager,

		pub.WriteUserChatMessage:            messageBaseUri + "/" + pub.WriteUserChatMessage,
		pub.WriteUserChatMessageAsManager:   messageBaseUri + "/" + pub.WriteUserChatMessageAsManager,
		pub.WriteUserChatMessageAsUser:      messageBaseUri + "/" + pub.WriteUserChatMessageAsUser,
		pub.WriteDirectChatMessageAsManager: messageBaseUri + "/" + pub.WriteDirectChatMessageAsManager,

		pvt.CreateDirectChat: directChatBaseUri + "/" + pvt.CreateDirectChat,

		pub.ManageUserChat: userChatBaseUri + "/" + pub.ManageUserChat,

		pub.GetManager:       managerBaseUri + "/" + pub.GetManager,
		pub.BatchGetManagers: managerBaseUri + "/" + pub.BatchGetManagers,
		pub.SearchManagers:   managerBaseUri + "/" + pub.SearchManagers,

		pub.GetUserChat: userChatBaseUri + "/" + pub.GetUserChat,

		pub.GetUser:    userBaseUri + "/" + pub.GetUser,
		pub.GetChannel: channelBaseUri + "/" + pub.GetChannel,
		pub.GetGroup:   groupBaseUri + "/" + pub.GetGroup,
		pvt.SearchUser: userBaseUri + "/" + pvt.SearchUser,

		pvt.FindOrCreateContactAndUser: mediumBaseUri + "/" + pvt.FindOrCreateContactAndUser,
		pvt.SearchUserChatsByContact:   userChatBaseUri + "/" + pvt.SearchUserChatsByContact,
		pvt.UpdateUserChatState:        userChatBaseUri + "/" + pvt.UpdateUserChatState,
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
