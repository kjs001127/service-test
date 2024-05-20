package dto_test

import (
	"testing"

	"github.com/channel-io/ch-app-store/api/http/account/dto"

	"github.com/stretchr/testify/assert"
)

func TestAppModifyValidate(t *testing.T) {
	t.Run("Test title length", func(t *testing.T) {
		req := dto.AppModifyRequest{
			Title: "정상적인 타이틀20글자보다 적음",
		}

		assert.Nil(t, req.Validate())
	})

	t.Run("Test under length 2", func(t *testing.T) {
		req := dto.AppModifyRequest{
			Title: "1",
		}

		assert.NotNil(t, req.Validate())
	})

	t.Run("Test over length 20", func(t *testing.T) {
		req := dto.AppModifyRequest{
			Title: "이건20글자가넘는타이틀이다한글로해도넘어야함동해물과백두산이",
		}

		assert.NotNil(t, req.Validate())
	})

	t.Run("Test description length", func(t *testing.T) {
		req := dto.AppModifyRequest{
			Title:       "정상적인 타이틀20글자보다 적음",
			Description: new(string),
		}

		assert.Nil(t, req.Validate())
	})

	t.Run("Test description over length 100", func(t *testing.T) {
		req := dto.AppModifyRequest{
			Title:       "정상적인 타이틀20글자보다 적음",
			Description: new(string),
		}
		*req.Description = "이건100글자가넘는타이틀이다한글로해도넘어야함동해물과백두산이이건100글자가넘는타이틀이다한글로해도넘어야함동해물과백두산이12312312312312af123afasdfasd3asdfasdfadfa"
		assert.NotNil(t, req.Validate())
	})
}
