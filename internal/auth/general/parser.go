package general

import (
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/channel-io/ch-proto/auth/v1/go/model"
)

type Token string

func (t Token) Value() string {
	return string(t)
}

type Parser interface {
	Parse(ctx context.Context, token string) (ParsedRBACToken, error)
}

type ParserImpl struct {
	roleCli       RoleFetcher
	jwtServiceKey string
}

func NewParser(roleCli RoleFetcher, jwtServiceKey string) *ParserImpl {
	return &ParserImpl{roleCli: roleCli, jwtServiceKey: jwtServiceKey}
}

func (f *ParserImpl) Parse(ctx context.Context, token string) (ParsedRBACToken, error) {

	parsedToken, err := f.parseJWT(token)
	if err != nil {
		return ParsedRBACToken{}, err
	}

	getRoleRes, err := f.roleCli.GetRole(ctx, parsedToken.RoleId)
	if err != nil {
		return ParsedRBACToken{}, err
	}

	ret, err := f.merge(getRoleRes.Role, parsedToken)
	if err != nil {
		return ParsedRBACToken{}, nil
	}

	ret.Type, ret.ID, _ = strings.Cut(parsedToken.Identity, "-")
	ret.Token = Token(token)

	return ret, nil
}

func (f *ParserImpl) parseJWT(token string) (*RBACToken, error) {
	parsed, err := jwt.ParseWithClaims(token, &RBACToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(f.jwtServiceKey), nil
		}
		return ParsedRBACToken{}, errors.New("signing method differs")
	})

	if err != nil {
		e, ok := err.(*jwt.ValidationError)
		if !ok || ok && e.Errors&jwt.ValidationErrorIssuedAt == 0 { // Don't report error that token used before issued.
			return &RBACToken{}, err
		}
	}

	return parsed.Claims.(*RBACToken), nil
}

func (f *ParserImpl) merge(role *model.Role, token *RBACToken) (ParsedRBACToken, error) {
	ret := ParsedRBACToken{
		Actions: make(map[Service][]Action),
		Scopes:  make(Scopes),
	}

	for _, scopeKeyVal := range token.Scope {
		key, val, ok := strings.Cut(scopeKeyVal, "-")
		if !ok {
			return ParsedRBACToken{}, errors.New("invalid scope")
		}
		ret.Scopes[key] = append(ret.Scopes[key], val)
	}

	for _, c := range role.Claims {
		ret.Actions[Service(c.Service)] = append(ret.Actions[Service(c.Service)], Action(c.Action))
	}

	return ret, nil
}

type RBACToken struct {
	jwt.StandardClaims
	RoleId   string   `json:"roleId"`
	Scope    []string `json:"scope"`
	Identity string   `json:"identity"`
}

type Service string
type Action string
type Scopes map[string][]string

const wildcard = "*"

type Caller struct {
	Type string
	ID   string
}

type ParsedRBACToken struct {
	Actions map[Service][]Action
	Scopes  Scopes
	Caller
	Token Token
}

func (p *ParsedRBACToken) GetCaller() Caller {
	return p.Caller
}

func (p *ParsedRBACToken) CheckAction(service Service, action Action) bool {
	if len(service) <= 0 || len(action) <= 0 {
		return false
	}

	if _, ok := p.Actions[wildcard]; ok {
		return true
	}

	actions := p.Actions[service]
	for _, a := range actions {
		if a == action || a == wildcard {
			return true
		}
	}
	return false
}

func (p *ParsedRBACToken) CheckScope(key string, value string) bool {
	if len(value) <= 0 || len(key) <= 0 {
		return false
	}

	if _, scopeExists := p.Scopes[key]; !scopeExists {
		return true
	}

	for _, s := range p.Scopes[key] {
		if s == value || s == wildcard {
			return true
		}
	}
	return false
}

func (p *ParsedRBACToken) CheckScopes(scopes Scopes) bool {
	for key, vals := range scopes {
		for _, val := range vals {
			if !p.CheckScope(key, val) {
				return false
			}
		}
	}
	return true
}

func (p *ParsedRBACToken) GetScope(scope string) []string {
	return p.Scopes[scope]
}
