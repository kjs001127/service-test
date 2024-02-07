package general

import (
	"context"
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt"
	"github.com/golang/protobuf/proto"

	"github.com/channel-io/ch-app-store/auth/appauth"
	"github.com/channel-io/ch-proto/auth/v1/go/service"
)

type JwtServiceKey string

type Token string

func (t Token) Value() string {
	return string(t)
}

func Header() string {
	return "x-access-token"
}

type Parser struct {
	cli            *resty.Client
	roleRequestUrl string
	jwtServiceKey  JwtServiceKey
}

func NewParser(cli *resty.Client, roleRequestUrl string, jwtServiceKey JwtServiceKey) *Parser {
	return &Parser{cli: cli, roleRequestUrl: roleRequestUrl, jwtServiceKey: jwtServiceKey}
}

func (f *Parser) Parse(ctx context.Context, token string) (ParsedRBACToken, error) {

	claims, err := f.parseJWT(token)
	if err != nil {
		return ParsedRBACToken{}, err
	}

	role, err := f.fetchRole(ctx, claims)
	if err != nil {
		return ParsedRBACToken{}, err
	}

	return f.merge(role, claims)
}

func (f *Parser) parseJWT(token string) (*RBACToken, error) {
	parsed, err := jwt.ParseWithClaims(token, &RBACToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return f.jwtServiceKey, nil
		}
		return ParsedRBACToken{}, errors.New("signing method differs")
	})
	if err != nil {
		return &RBACToken{}, err
	}
	if !parsed.Valid {
		return &RBACToken{}, errors.New("token invalid")
	}

	return parsed.Claims.(*RBACToken), nil
}

func (f *Parser) fetchRole(ctx context.Context, claims *RBACToken) (*service.GetRoleResult, error) {
	r := f.cli.R()
	r.SetContext(ctx)

	body, err := proto.Marshal(&service.GetRoleRequest{RoleId: claims.RoleId})
	if err != nil {
		return &service.GetRoleResult{}, err
	}
	r.SetBody(body)

	rawRes, err := r.Post(r.URL + "/v1/roles/getRole")
	if err != nil {
		return &service.GetRoleResult{}, err
	}

	var res service.GetRoleResult
	if err := proto.Unmarshal(rawRes.Body(), &res); err != nil {
		return &service.GetRoleResult{}, err
	}

	return &res, nil
}

func (f *Parser) merge(role *service.GetRoleResult, claims *RBACToken) (ParsedRBACToken, error) {
	ret := ParsedRBACToken{
		Actions:     make(map[Service][]Action),
		Authorities: make(appauth.Authorities),
	}

	for _, scopeKeyVal := range claims.Scope {
		key, val, ok := strings.Cut(scopeKeyVal, "-")
		if !ok {
			return ParsedRBACToken{}, errors.New("invalid scope")
		}
		ret.Authorities[key] = append(ret.Authorities[key], val)
	}

	for _, c := range role.Role.Claims {
		ret.Actions[Service(c.Service)] = append(ret.Actions[Service(c.Service)], Action(c.Action))
	}

	return ret, nil
}

type RBACToken struct {
	jwt.StandardClaims
	RoleId string   `json:"roleId"`
	Scope  []string `json:"scope"`
}

type Service string
type Action string

type ParsedRBACToken struct {
	Actions     map[Service][]Action
	Authorities appauth.Authorities
}

func (p *ParsedRBACToken) CheckAction(service Service, action Action) bool {
	actions := p.Actions[service]
	for _, a := range actions {
		if a == action || a == appauth.Wildcard {
			return true
		}
	}
	return false
}

func (p *ParsedRBACToken) CheckAuthority(key string, value string) bool {
	if len(key) <= 0 {
		return true
	}
	if _, ok := p.Authorities[appauth.Wildcard]; ok {
		return true
	}

	scopes := p.Authorities[key]
	for _, s := range scopes {
		if s == value || s == appauth.Wildcard {
			return true
		}
	}
	return false
}

func (p *ParsedRBACToken) CheckAuthorities(scopes appauth.Authorities) bool {
	for key, vals := range scopes {
		for _, val := range vals {
			if !p.CheckAuthority(key, val) {
				return false
			}
		}
	}
	return true
}
