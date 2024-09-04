package middleware

import (
	"log"
	"strconv"

	"github.com/channel-io/ch-app-store/lib/ratelimiter"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/gin-gonic/gin"
)

const (
	XRateLimitLimit           = "X-Rate-Limit-Limit"
	XRateLimitRemaining       = "X-Rate-Limit-Remaining"
	XRateLimitReset           = "X-Rate-Limit-Reset"
	XRateLimitWillBeThrottled = "X-Rate-Limit-Will-Be-Throttled"
)

type RateLimiter struct {
	client ratelimiter.Client
}

func NewRateLimiter(rateLimiterClient ratelimiter.Client) *RateLimiter {
	return &RateLimiter{
		client: rateLimiterClient,
	}
}

func (r *RateLimiter) Priority() int {
	return 5
}

func (r *RateLimiter) Handle(ctx *gin.Context) {
	key, exists := ctx.Get(ResolverKey)

	if !exists {
		ctx.Next()
	}

	resolved, ok := key.(Key)
	if !ok {
		return
	}

	req := ratelimiter.RateLimitRequest{
		RuleIDs:    []string{resolved.RuleID},
		Attributes: resolved.Attributes,
	}

	res, err := r.client.Acquire(ctx, req)
	if apierr.IsTooManyRequest(err) {

		result := res.Results[0]

		rateLimitErr := NewRateLimitError(result.Limit, result.Remaining, int(result.ResetAt.Unix()), true)

		ctx.Abort()
		_ = ctx.Error(rateLimitErr)
		return
	} else if err != nil {
		log.Printf("rate limit error occurred %v", err)
		return
	}

	// if rate limit is not exceeded, set rate limit headers
	rateLimitResult := res.Results[0]

	ctx.Header(XRateLimitRemaining, strconv.Itoa(rateLimitResult.Remaining))
	ctx.Header(XRateLimitLimit, strconv.Itoa(rateLimitResult.Limit))
	ctx.Header(XRateLimitReset, strconv.FormatInt(rateLimitResult.ResetAt.Unix(), 10))
}

type RateLimitError struct {
	apierr.HTTPErrorBuildable

	XRateLimitLimit           int
	XRateLimitRemaining       int
	XRateLimitReset           int
	XRateLimitWillBeThrottled bool
}

func (r *RateLimitError) HTTPStatusCode() int {
	return 429
}

func (r *RateLimitError) ErrorName() string {
	return "RateLimitError"
}

func (r *RateLimitError) Causes() []*apierr.Cause {
	return []*apierr.Cause{
		{
			Message: "Rate limit exceeded",
			Detail: map[string]interface{}{
				XRateLimitLimit:           r.XRateLimitLimit,
				XRateLimitRemaining:       r.XRateLimitRemaining,
				XRateLimitReset:           r.XRateLimitReset,
				XRateLimitWillBeThrottled: r.XRateLimitWillBeThrottled,
			},
		},
	}
}

func (r *RateLimitError) Error() string {
	return r.ErrorName()
}

func NewRateLimitError(limit, remaining, reset int, willBeThrottled bool) *RateLimitError {
	return &RateLimitError{
		XRateLimitLimit:           limit,
		XRateLimitRemaining:       remaining,
		XRateLimitReset:           reset,
		XRateLimitWillBeThrottled: willBeThrottled,
	}
}
