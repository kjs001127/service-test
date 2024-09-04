package ratelimiter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	Acquire(ctx context.Context, req RateLimitRequest) (RateLimitResponse, error)
}

type client struct {
	cli            *resty.Client
	rateLimiterURL string
}

type RateLimitRequest struct {
	RuleIDs    []string
	Attributes map[string]string
}

type RateLimitResult struct {
	RuleID      string
	Pass        bool
	ResourceKey string
	Limit       int
	Remaining   int
	ResetAt     time.Time
}

type RateLimitResponse struct {
	Results []RateLimitResult
}

func NewRateLimiterClient(cli *resty.Client, rateLimiterURL string) Client {
	return &client{cli: cli, rateLimiterURL: rateLimiterURL}
}

func (c client) Acquire(ctx context.Context, req RateLimitRequest) (RateLimitResponse, error) {
	r := c.cli.R()
	r.SetContext(ctx)

	r.
		SetBody(req).
		SetHeader("Content-Type", "application/json")

	resp, err := r.Post(c.rateLimiterURL + "/v1/rate-limit")

	var unmarshalled RateLimitResponse

	if err != nil {
		return RateLimitResponse{}, err
	}

	if !resp.IsSuccess() {
		if resp.StatusCode() == 429 {
			if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
				return RateLimitResponse{}, err
			}
			return unmarshalled, apierr.TooManyRequest(fmt.Errorf("rate limit failed, response: %s", resp.Body()))
		} else {
			return RateLimitResponse{}, apierr.UnprocessableEntity(fmt.Errorf("rate limit failed, response: %s", resp.Body()))
		}
	}

	if err := json.Unmarshal(resp.Body(), &unmarshalled); err != nil {
		return RateLimitResponse{}, err
	}
	return unmarshalled, nil
}
