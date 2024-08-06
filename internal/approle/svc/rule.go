package svc

import (
	protomodel "github.com/channel-io/ch-proto/auth/v1/go/model"
)

type ClaimsFactory func(string) []*protomodel.Claim

type TypeRule struct {
	AvailableClaimsOf ClaimsFactory
	GrantTypes        []protomodel.GrantType
	PrincipalTypes    []string
	DefaultClaimsOf   ClaimsFactory
}
