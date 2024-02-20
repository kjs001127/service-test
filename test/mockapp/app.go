package mockapp

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"io"

	"github.com/channel-io/go-lib/pkg/errors/apierr"
	"github.com/channel-io/go-lib/pkg/ptr"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
	cmd "github.com/channel-io/ch-app-store/internal/command/domain"
)

func SetUpMockApp(manager app.AppManager, reg *cmd.RegisterSvc) error {
	_, err := manager.Create(context.Background(), &app.App{
		ID:        "1",
		State:     app.AppStateStable,
		Title:     "TestApp",
		IsPrivate: true,
	})
	if err != nil {
		return err
	}

	_ = reg.Register(context.Background(), "1",
		[]*cmd.Command{
			{
				ID:                       "1",
				Name:                     "deskCommand",
				Description:              ptr.Of("데스크 테스트 커맨드 입니다"),
				Scope:                    cmd.ScopeDesk,
				AutoCompleteFunctionName: ptr.Of("deskAutoComplete"),
				ParamDefinitions: []*cmd.ParamDefinition{
					{
						Name:         "파라미터1",
						Description:  "파라미터 설명입니다",
						Type:         cmd.ParamTypeString,
						Required:     true,
						AutoComplete: true,
					},
				},
				ActionFunctionName: "deskActionFunction",
			},
			{
				ID:          "2",
				Name:        "frontCommand",
				Description: ptr.Of("프론트 테스트 커맨드 입니다"),
				Scope:       cmd.ScopeFront,
				ParamDefinitions: []*cmd.ParamDefinition{
					{
						Name:        "파라미터1",
						Description: "파라미터 설명입니다",
						Type:        cmd.ParamTypeString,
						Required:    true,
						Choices: cmd.Choices{
							{Name: "테스트1", Value: "테스트값1"},
							{Name: "테스트2", Value: "테스트값2"},
						},
						AutoComplete: false,
					},
				},
				ActionFunctionName: "frontActionFunction",
			},
		},
	)
	return nil
}

type InstallHandler struct {
}

func (i InstallHandler) OnInstall(ctx context.Context, app *app.App, channelID string) error {
	return nil
}

func (i InstallHandler) OnUnInstall(ctx context.Context, app *app.App, channelID string) error {
	return nil
}

type InvokeHandler struct {
}

func (i InvokeHandler) Invoke(ctx context.Context, target *app.App, request app.JsonFunctionRequest) app.JsonFunctionResponse {
	if target.ID != "1" {
		return app.WrapErr(errors.New("cannot find app"))
	}

	var ret any
	switch request.Method {
	case "deskActionFunction":
		ret = cmd.Action{
			Type: "wam",
			Attributes: map[string]any{
				"name": "deskWam",
				"context": map[string]any{
					"testParam": "value",
				},
			},
		}
	case "frontActionFunction":
		ret = cmd.Action{
			Type: "wam",
			Attributes: map[string]any{
				"name": "frontWam",
				"context": map[string]any{
					"testParam": "value",
				},
			},
		}
	case "deskAutoComplete":
		ret = cmd.Choices{
			{Name: "테스트0", Value: "testValue0"},
			{Name: "테스트1", Value: "testValue1"},
		}
	default:
		return app.WrapErr(apierr.NotFound(errors.New("no command found")))
	}

	res, _ := json.Marshal(ret)
	return app.JsonFunctionResponse{
		Result: res,
	}
}

// go:embed resources/*
var fs embed.FS

type FileStreamer struct {
}

func (f FileStreamer) StreamFile(ctx context.Context, appID string, path string, writer io.Writer) error {
	bytes, err := fs.ReadFile(path)
	if err != nil {
		return err
	}

	if _, err = writer.Write(bytes); err != nil {
		return err
	}

	return nil
}

type ConfigValidator struct {
}

func (c ConfigValidator) ValidateConfig(ctx context.Context, app *app.App, channelID string, input app.ConfigMap) error {
	return nil
}
