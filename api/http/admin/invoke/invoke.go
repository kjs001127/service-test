package invoke

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/channel-io/ch-app-store/api/http/shared/dto"
	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

const (
	callerAdmin = "admin"
	idInferred  = "-"
)

// invoke godoc
//
//	@Summary	invoke Function
//	@Tags		Admin
//
//	@Param		appID				path		string				true	"id of App to invoke Function"
//	@Param		name				path		string				true	"name of Function to invoke"
//	@Param		dto.JsonRPCRequest	body		dto.JsonRPCRequest	true	"body of Function to invoke"
//
//	@Success	200					{object}	json.RawMessage
//	@Router		/admin/channels/{channelID}/apps/{appID}/functions/{name} [post]
func (h *Handler) invoke(ctx *gin.Context) {
	appID, name, channelID := ctx.Param("id"), ctx.Param("name"), ctx.Param("channelID")

	var req dto.JsonRPCRequest
	if err := ctx.ShouldBindBodyWith(req, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.HttpBadRequestError(err))
		return
	}

	res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest[json.RawMessage]{
		Endpoint: app.Endpoint{
			AppID:        appID,
			FunctionName: name,
		},
		Body: app.Body[json.RawMessage]{
			Caller: app.Caller{
				Type: callerAdmin,
				ID:   idInferred,
			},
			Params: req.Params,
		},
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) brief(ctx *gin.Context) {

	channelID := ctx.Param("channelID")

	installed, err := h.querySvc.QueryAll(ctx, channelID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	briefs, err := h.briefRepo.FetchAll(ctx, app.AppIDsOf(installed.AppChannels))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.HttpUnprocessableEntityError(err))
		return
	}

	ch := make(chan json.RawMessage, len(briefs))
	var wg sync.WaitGroup
	wg.Add(len(briefs))

	for _, brief := range briefs {
		brief := brief
		go func() {
			res, err := h.invoker.InvokeChannelFunction(ctx, channelID, app.FunctionRequest[json.RawMessage]{
				Endpoint: app.Endpoint{
					AppID:        brief.AppID,
					FunctionName: brief.BriefFunctionName,
				},
				Body: app.Body[json.RawMessage]{
					Caller: app.Caller{
						Type: callerAdmin,
						ID:   idInferred,
					},
				},
			})
			if err != nil {
				ch <- res
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var ret []json.RawMessage
	for res := range ch {
		ret = append(ret, res)
	}

	ctx.JSON(http.StatusOK, ret)
}
