package general

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/proto"

	"github.com/channel-io/ch-app-store/internal/remoteapp/domain"
	"github.com/channel-io/ch-proto/auth/v1/go/model"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

const (
	version     = "v1"
	roleBaseUri = "/admin/auth/" + version + "/roles"
	fetchRole   = roleBaseUri + "/getRole"
	createRole  = roleBaseUri + "/createRole"
	deleteRole  = roleBaseUri + "/deleteRole"
)

type RoleClient struct {
	cli     *resty.Client
	authUrl string
}

func NewRoleClient(cli *resty.Client, authUrl string) *RoleClient {
	return &RoleClient{cli: cli, authUrl: authUrl}
}

func (f *RoleClient) FetchRole(ctx context.Context, roleID string) (*service.GetRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)

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

func (f *RoleClient) CreateRole(ctx context.Context, claims []*model.Claim) (*service.CreateRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)

	body, err := proto.Marshal(&service.CreateRoleRequest{Claims: claims})
	if err != nil {
		return nil, err
	}
	r.SetBody(body)
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

type RoleClientAdapter struct {
	cli *RoleClient
}

func (r RoleClientAdapter) ReadRole(ctx context.Context, roleID string) (*domain.Role, error) {
	role, err := r.cli.FetchRole(ctx, roleID)
	if err != nil {
		return nil, err
	}
	return marshal(role.Role), nil
}

func (r RoleClientAdapter) CreateRole(ctx context.Context, request *domain.Role) (domain.RoleWithCredential, error) {
	created, err := r.cli.CreateRole(ctx, unmarshalClaims(request.Claims))
	if err != nil {
		return domain.RoleWithCredential{}, err
	}
	return domain.RoleWithCredential{
		Role: marshal(created.Role),
		RoleCredentials: domain.RoleCredentials{
			ClientID: created.Credentials.ClientId,
			Secret:   created.Credentials.ClientSecret,
		},
	}, nil
}

func (r RoleClientAdapter) UpdateRole(ctx context.Context, roleID string, claims []domain.Claim) (*domain.Role, error) {
	return nil, errors.New("not implemented")
}

func (r RoleClientAdapter) DeleteRole(ctx context.Context, roleID string) error {
	_, err := r.cli.DeleteRole(ctx, roleID)
	if err != nil {
		return err
	}
	return nil
}

func NewRoleClientAdapter(cli *RoleClient) *RoleClientAdapter {
	return &RoleClientAdapter{cli: cli}
}

func marshal(role *model.Role) *domain.Role {
	return &domain.Role{
		ID:     role.Id,
		Claims: marshalClaims(role.Claims),
	}
}
func marshalClaims(claims []*model.Claim) []domain.Claim {
	ret := make([]domain.Claim, 0, len(claims))
	for _, c := range claims {
		ret = append(ret, domain.Claim{
			Service: c.Service,
			Action:  c.Action,
			Scopes:  c.Scope,
		})
	}
	return ret
}
func unmarshalClaims(claims []domain.Claim) []*model.Claim {
	ret := make([]*model.Claim, 0, len(claims))
	for _, c := range claims {
		ret = append(ret, &model.Claim{
			Service: c.Service,
			Action:  c.Action,
			Scope:   c.Scopes,
		})
	}
	return ret
}
