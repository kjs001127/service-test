package middleware

import (
	"log"
	"strings"

	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/lib/ratelimiter"

	"github.com/gin-gonic/gin"
)

type XSessionTokenKeyResolver struct {
	ruleSet *ratelimiter.RuleSet
}

func NewXSessionKeyResolver(ruleSet *ratelimiter.RuleSet) *XSessionTokenKeyResolver {
	return &XSessionTokenKeyResolver{
		ruleSet: ruleSet,
	}
}

func (m *XSessionTokenKeyResolver) Priority() int {
	return 4
}

func (m *XSessionTokenKeyResolver) Handle(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.RequestURI, FrontPathPrefix) {
		return
	}

	user := User(ctx)

	attributes := map[string]string{
		"callerType": UserKey,
		"callerId":   user.ID,
	}

	ruleID, err := m.ruleSet.GetRule(ratelimiter.UserRuleKey)
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
