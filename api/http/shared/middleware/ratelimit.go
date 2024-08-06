package middleware

import (
	"log"
	"net/http"
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
		log.Printf("error occurred while resolving rate limit key")
		return
	}

	req := ratelimiter.RateLimitRequest{
		RuleIDs:    []string{resolved.RuleID},
		Attributes: resolved.Attributes,
	}

	res, err := r.client.Acquire(ctx, req)
	if apierr.IsTooManyRequest(err) {

		result := res.Results[0]

		rateLimitResponse := make(map[string]interface{})
		rateLimitResponse[XRateLimitLimit] = result.Limit
		rateLimitResponse[XRateLimitRemaining] = result.Remaining
		rateLimitResponse[XRateLimitReset] = result.ResetAt
		rateLimitResponse[XRateLimitWillBeThrottled] = true

		ctx.JSON(http.StatusTooManyRequests, rateLimitResponse)
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
