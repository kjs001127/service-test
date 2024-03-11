package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/channel-io/ch-app-store/lib/log"
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
	} else if ctx.Writer.Status() >= http.StatusBadRequest {
		_ = ctx.Error(apierr.BadRequest(errors.New("unknown bad request error")))
	}
}

func (l *ErrHandler) log(ctx *gin.Context, dto *errorsDTO) {
	body, _ := io.ReadAll(ctx.Request.Body)
	if dto.Status >= http.StatusInternalServerError {
		l.logger.Errorw(ctx, "http request failed",
			"uri", ctx.Request.RequestURI,
			"status", ctx.Writer.Status(),
			"body", json.RawMessage(body),
			"err", dto.Errors,
		)
	} else if dto.Status >= http.StatusBadRequest {
		l.logger.Warnw(ctx, "http request failed",
			"uri", ctx.Request.RequestURI,
			"status", ctx.Writer.Status(),
			"body", json.RawMessage(body),
			"err", dto.Errors,
		)
	}
}

func errorsDTOFrom(ctx *gin.Context) *errorsDTO {
	var httpErrorBuildable apierr.HTTPErrorBuildable
	for _, err := range ctx.Errors {
		if errors.As(err.Unwrap(), &httpErrorBuildable) {
			dto := newErrorsDTOFromHTTPErrorBuildable(httpErrorBuildable)
			return dto
		}
	}
	return newErrorsDTO(http.StatusUnprocessableEntity)
}

type errorsDTO struct {
	Status   int        `json:"status"`
	Type     string     `json:"type"`
	Language string     `json:"language"`
	Errors   []errorDTO `json:"errors"`
}

type errorDTO map[string]any

func newErrorsDTO(statusCode int, causes ...*apierr.Cause) *errorsDTO {
	var errDTOs []errorDTO
	for _, cause := range causes {
		errDTOs = append(errDTOs, newErrorDTO(cause))
	}

	return &errorsDTO{
		Status:   statusCode,
		Type:     http.StatusText(statusCode),
		Language: "ko", // TODO(billo): 언어 지원되면 DTO 고치기
		Errors:   errDTOs,
	}
}

func newErrorsDTOFromHTTPErrorBuildable(buildable apierr.HTTPErrorBuildable) *errorsDTO {
	return newErrorsDTO(
		buildable.HTTPStatusCode(),
		buildable.Causes()...,
	)
}

func newErrorDTO(cause *apierr.Cause) errorDTO {
	dto := make(map[string]any)
	dto["message"] = cause.Message
	for k, v := range cause.Detail {
		dto[k] = v
	}

	return dto
}
