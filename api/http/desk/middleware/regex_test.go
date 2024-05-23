package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexMatch(t *testing.T) {
	urls := []string{
		"/desk/v1/channels/123",
		"/desk/v1/channels/123/apps",
		"/desk/v12/channels/123/apps/123/commands",
	}

	for _, url := range urls {
		assert.True(t, regex.MatchString(url))
	}
}

func TestRegexUnmatch(t *testing.T) {
	urls := []string{
		"/front/v1/channels/123",
		"/desk/account/channels/123/apps",
		"/desk/account/channels",
	}
	for _, url := range urls {
		assert.False(t, regex.MatchString(url))
	}
}
