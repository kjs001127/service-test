package general

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/proto"

	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

const (
	version     = "v1"
	roleBaseUri = "/api/admin/auth/" + version + "/roles"
	fetchRole   = roleBaseUri + "/getRole"
	createRole  = roleBaseUri + "/createRole"
	deleteRole  = roleBaseUri + "/deleteRole"
)

type RoleClient struct {
	cli     *resty.Client
	authUrl string
}

func (f *RoleClient) UpdateRole(ctx context.Context, roleID string, request *service.ReplaceRoleClaimsRequest) (*service.ReplaceRoleClaimsResult, error) {
	//TODO implement me
	panic("implement me")
}

func NewRoleClient(cli *resty.Client, authUrl string) *RoleClient {
	return &RoleClient{cli: cli, authUrl: authUrl}
}

func (f *RoleClient) GetRole(ctx context.Context, roleID string) (*service.GetRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)
	r.SetHeader("Content-Type", "application/x-protobuf")
	r.SetHeader("Accept", "application/x-protobuf")

	body, err := proto.Marshal(&service.GetRoleRequest{RoleId: roleID})
	if err != nil {
		return &service.GetRoleResult{}, err
	}
	r.SetBody(body)

	rawRes, err := r.Post(f.authUrl + fetchRole)
	if err != nil {
		return &service.GetRoleResult{}, err
	}

	var res service.GetRoleResult
	if err := proto.Unmarshal(rawRes.Body(), &res); err != nil {
		return &service.GetRoleResult{}, err
	}

	return &res, nil
}

func (f *RoleClient) CreateRole(ctx context.Context, request *service.CreateRoleRequest) (*service.CreateRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)

	body, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	r.SetBody(body)
	r.SetHeader("Content-Type", "application/x-protobuf")
	r.SetHeader("Accept", "application/x-protobuf")

	rawRes, err := r.Post(f.authUrl + createRole)
	if err != nil {
		return nil, err
	}
	var res service.CreateRoleResult
	if err := proto.Unmarshal(rawRes.Body(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (f *RoleClient) DeleteRole(ctx context.Context, roleID string) (*service.DeleteRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)

	body, err := proto.Marshal(&service.DeleteRoleRequest{RoleId: roleID})
	if err != nil {
		return nil, err
	}
	r.SetBody(body)
	r.SetHeader("Content-Type", "application/x-protobuf")
	r.SetHeader("Accept", "application/x-protobuf")

	rawRes, err := r.Post(f.authUrl + deleteRole)
	if err != nil {
		return nil, err
	}

	var res service.DeleteRoleResult
	if err := proto.Unmarshal(rawRes.Body(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
