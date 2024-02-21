package function

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/proto"

	"github.com/channel-io/ch-app-store/internal/native/domain"
	"github.com/channel-io/ch-proto/coreapi/v1/go/service"
)

const baseUri = "/api/admin/core"

type CoreApi struct {
	adminUrl string
	resty    *resty.Client
}

func (a *CoreApi) Provide() domain.NativeFunctions {
	return domain.NativeFunctions{
		"writeUserChatMessage": a.WriteUserChatMessage,
	}
}

func NewCoreApi(adminUrl string, resty *resty.Client) *CoreApi {
	return &CoreApi{adminUrl: adminUrl, resty: resty}
}

func (a *CoreApi) WriteUserChatMessage(ctx context.Context, params json.RawMessage, token domain.Token) domain.NativeFunctionResponse {
	var req service.WriteUserChatMessageRequest
	if err := json.Unmarshal(params, &req); err != nil {
		domain.WrapError(err)
	}
	marshaledReq, err := proto.Marshal(&req)
	if err != nil {
		return domain.WrapError(err)
	}

	r := a.resty.R()
	r.SetHeader(token.Type, token.Value)
	r.SetHeader("Content-Type", "application/x-protobuf")
	r.SetBody(marshaledReq)
	r.SetContext(ctx)
	resp, err := r.Post(a.adminUrl + baseUri + "/writeUserChatMessage")
	if err != nil {
		return domain.WrapError(err)
	}

	var res service.WriteUserChatMessageResult
	if err := proto.Unmarshal(resp.Body(), &res); err != nil {
		return domain.WrapError(err)
	}

	if resp.IsError() {
		return domain.WrapError(fmt.Errorf("request failed, body: %s", resp.Body()))
	}

	marshaledRes, err := json.Marshal(&res)
	if err != nil {
		return domain.WrapError(err)
	}

	return domain.NativeFunctionResponse{
		Result: marshaledRes,
	}
}
