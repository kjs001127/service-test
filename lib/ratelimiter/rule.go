package ratelimiter

import (
	"github.com/channel-io/ch-app-store/config"
	"github.com/pkg/errors"
)

const (
	ManagerRuleKey = "manager"
	AppRuleKey     = "app"
	UserRuleKey    = "user"
)

type Key string

func Convert(param string) Key {
	return Key(param)
}

type Rules map[string]string

func (r Rules) AddRule(key, value string) {
	r[key] = value
}

func NewRules() Rules {
	return make(Rules)
}

type RuleSet struct {
	rules Rules
}

func NewRuleSet() *RuleSet {
	cfg := config.Get()
	rules := NewRules()

	rules.AddRule(ManagerRuleKey, cfg.RateLimiter.ManagerRuleID)
	rules.AddRule(UserRuleKey, cfg.RateLimiter.UserRuleID)
	rules.AddRule(AppRuleKey, cfg.RateLimiter.AppRuleID)

	return &RuleSet{
		rules: rules,
	}
}

func (r *RuleSet) GetRule(key Key) (string, error) {
	rawKey := string(key)
	rule, ok := r.rules[rawKey]
	if !ok {
		return "", errors.New("rule not found")
	}
	return rule, nil
}
