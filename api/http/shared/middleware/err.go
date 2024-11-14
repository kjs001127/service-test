package middleware

import (
	"net/http"

	"github.com/channel-io/ch-app-store/internal/shared/principal"
	"github.com/channel-io/ch-app-store/lib/i18n"
	"github.com/channel-io/ch-app-store/lib/log"

	"github.com/channel-io/go-lib/pkg/errors/apierr"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type ErrHandler struct {
	logger log.ContextAwareLogger
}

func (l *ErrHandler) Priority() int {
	return -1
}

func NewErrHandler(logger log.ContextAwareLogger) *ErrHandler {
	return &ErrHandler{logger: logger}
}

func (l *ErrHandler) Handle(ctx *gin.Context) {
	ctx.Next()

	l.propagateStatusToErr(ctx)

	if len(ctx.Errors) <= 0 {
		return
	}

	dto := errorsDTOFrom(ctx)
	l.log(ctx, dto)
	ctx.JSON(dto.Status, dto)
}

func (l *ErrHandler) propagateStatusToErr(ctx *gin.Context) {
	if len(ctx.Errors) > 0 {
		return
	}

	if ctx.Writer.Status() >= http.StatusInternalServerError {
		_ = ctx.Error(errors.New("unknown internal error"))
	} else if ctx.Writer.Status() == http.StatusTooManyRequests {
		return
	} else if ctx.Writer.Status() >= http.StatusBadRequest {
		_ = ctx.Error(apierr.BadRequest(errors.New("unknown bad request error")))
	}
}

func (l *ErrHandler) log(ctx *gin.Context, dto *errorsDTO) {
	if dto.Status >= http.StatusInternalServerError {
		l.logger.Errorw(ctx, "http request failed",
			"uri", ctx.Request.RequestURI,
			"status", ctx.Writer.Status(),
			"err", dto.Errors,
		)
	} else if dto.Status >= http.StatusBadRequest {
		l.logger.Warnw(ctx, "http request failed",
			"uri", ctx.Request.RequestURI,
			"status", ctx.Writer.Status(),
			"err", dto.Errors,
		)
	}
}

func errorsDTOFrom(ctx *gin.Context) *errorsDTO {
	var httpErrorBuildable apierr.HTTPErrorBuildable
	for _, err := range ctx.Errors {
		if errors.As(err.Unwrap(), &httpErrorBuildable) {
			dto := newErrorsDTOFromHTTPErrorBuildable(ctx, httpErrorBuildable)
			return dto
		} else {
			return newErrorsDTO(ctx, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity, apierr.NewCause(err))
		}
	}
	return newErrorsDTO(ctx, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type errorsDTO struct {
	Status   int        `json:"status"`
	Type     string     `json:"type"`
	Language string     `json:"language"`
	Errors   []errorDTO `json:"errors"`
}

type errorDTO map[string]any

func newErrorsDTO(ctx *gin.Context, errType string, statusCode int, causes ...*apierr.Cause) *errorsDTO {
	var errDTOs []errorDTO
	for _, cause := range causes {
		errDTOs = append(errDTOs, newErrorDTO(ctx, cause))
	}

	lang := i18n.DefaultLanguage.String()
	rawRequester, _ := ctx.Get(RequesterKey)
	requester := rawRequester.(principal.Requester)
	if i18n.IsValid(requester.Language) {
		lang = requester.Language
	}

	return &errorsDTO{
		Status:   statusCode,
		Type:     errType,
		Language: lang,
		Errors:   errDTOs,
	}
}

func newErrorsDTOFromHTTPErrorBuildable(ctx *gin.Context, buildable apierr.HTTPErrorBuildable) *errorsDTO {
	return newErrorsDTO(ctx, buildable.ErrorName(), buildable.HTTPStatusCode(), buildable.Causes()...)
}

func newErrorDTO(ctx *gin.Context, cause *apierr.Cause) errorDTO {
	dto := make(map[string]any)
	dto["message"] = i18n.Translate(ctx, cause.Message)
	for k, v := range cause.Detail {
		dto[k] = v
	}

	return dto
}
