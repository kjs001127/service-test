package middleware

import (
	"github.com/channel-io/ch-app-store/internal/auth/principal"
	"github.com/channel-io/ch-app-store/lib/i18n"

	"github.com/gin-gonic/gin"
	languageUtil "golang.org/x/text/language"
)

const (
	AcceptLanguage = "Accept-Language"
	RequesterKey   = "Requester"
)

type Request struct {
	i18n i18n.I18n
}

func NewRequest(i18n i18n.I18n) *Request {
	return &Request{i18n: i18n}
}

func (r *Request) Priority() int {
	return 1
}

func (r *Request) Handle(ctx *gin.Context) {
	var language string

	tags, _, err := languageUtil.ParseAcceptLanguage(ctx.GetHeader(AcceptLanguage))
	if err != nil {
		language = i18n.DefaultLanguage.String()
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
	ctx.Set(i18n.LocalizerKey, r.i18n.GetLocalizer(language))
}

func Requester(ctx *gin.Context) principal.Requester {
	rawRequester, _ := ctx.Get(RequesterKey)

	return rawRequester.(principal.Requester)
}
