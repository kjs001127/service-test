package middleware

import (
	"log"
	"strings"

	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/lib/ratelimiter"

	"github.com/gin-gonic/gin"
)

type XAccessTokenKeyResolver struct {
	ruleSet *ratelimiter.RuleSet
}

func NewRbacKeyResolver(ruleSet *ratelimiter.RuleSet) *XAccessTokenKeyResolver {
	return &XAccessTokenKeyResolver{
		ruleSet: ruleSet,
	}
}

func (m *XAccessTokenKeyResolver) Priority() int {
	return 4
}

func (m *XAccessTokenKeyResolver) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, generalPathPrefix) {
		return
	}

	rbac, exists := Rbac(ctx)
	if !exists {
		return
	}

	attributes := map[string]string{
		"callerType": rbac.Caller.Type,
		"callerId":   rbac.Caller.ID,
	}

	converted := ratelimiter.Convert(rbac.Caller.Type)

	ruleID, err := m.ruleSet.GetRule(converted)
	if err != nil {
		log.Println(err)
		return
	}

	key := middleware.Key{
		RuleID:     ruleID,
		Attributes: attributes,
	}

	ctx.Set(middleware.ResolverKey, key)
}
