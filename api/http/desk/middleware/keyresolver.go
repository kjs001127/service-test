package middleware

import (
	"github.com/channel-io/ch-app-store/api/http/shared/middleware"
	"github.com/channel-io/ch-app-store/lib/ratelimiter"
	"log"

	"github.com/gin-gonic/gin"
)

type XAccountKeyResolver struct {
	ruleSet *ratelimiter.RuleSet
}

func NewXAccountKeyResolver(ruleSet *ratelimiter.RuleSet) *XAccountKeyResolver {
	return &XAccountKeyResolver{
		ruleSet: ruleSet,
	}
}

func (m *XAccountKeyResolver) Priority() int {
	return 4
}

func (m *XAccountKeyResolver) Handle(ctx *gin.Context) {
	if !regex.MatchString(ctx.Request.RequestURI) {
		return
	}

	manager := Manager(ctx)

	attributes := map[string]string{
		"callerType": "manager",
		"callerId":   manager.ID,
	}

	ruleID, err := m.ruleSet.GetRule(ratelimiter.ManagerRuleKey)
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
