package middleware

import (
	"github.com/channel-io/ch-app-store/internal/auth/principal"

	"github.com/gin-gonic/gin"
	languageUtil "golang.org/x/text/language"
)

const (
	AcceptLanguage = "Accept-Language"
	RequesterKey   = "Requester"
)

type Request struct{}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) Priority() int {
	return 1
}

func (r *Request) Handle(ctx *gin.Context) {
	var language string

	tags, _, err := languageUtil.ParseAcceptLanguage(ctx.GetHeader(AcceptLanguage))
	if err != nil {
		language = "en"
	}

	for _, tag := range tags {
		lang, conf := tag.Base()
		if conf >= languageUtil.High {
			language = lang.String()
			break
		}
	}

	requester := principal.Requester{
		Language: language,
	}

	ctx.Set(RequesterKey, requester)
}

func Requester(ctx *gin.Context) principal.Requester {
	rawRequester, _ := ctx.Get(RequesterKey)

	return rawRequester.(principal.Requester)
}
