package middleware

const (
	ResolverKey = "RateLimitKeyResolver"
)

type Key struct {
	RuleID     string
	Attributes map[string]string
}
