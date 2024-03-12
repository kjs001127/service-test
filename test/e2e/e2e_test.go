package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	app "github.com/channel-io/ch-app-store/internal/app/svc"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

const srvPort = 3020

func Test(t *testing.T) {
	tests := []testInfo{
		{
			name: "ping test",
			req: request{
				method: "GET",
				path:   "/ping",
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "create remote app in admin",
			req: request{
				method: "POST",
				path:   "/admin/apps",
				body: map[string]any{
					"title":         "TestApp",
					"isPrivate":     false,
					"configSchemas": []string{},
					"roles":         []string{},
					"wamUrl":        FakeWamUrl,
					"functionUrl":   FakeFunctionUrl,
				},
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusCreated,
				assertionFunc: func(t *testing.T, resMap map[string]any) {
					assert.Equal(t, "TestApp", resMap["title"])
					assert.Equal(t, false, resMap["isPrivate"])
					assert.Equal(t, FakeFunctionUrl, resMap["functionUrl"])
					assert.Equal(t, FakeWamUrl, resMap["wamUrl"])
					assert.Equal(t, nil, resMap["roles"]) // omit empty
				},
			},
		},
		{
			name: "register command to app in admin",
			req: request{
				method: "POST",
				path:   "/admin/apps/{appId}/commands",
				body: toJSONMap(dto.RegisterRequest{Commands: []*model.Command{
					{
						Name:               "testCommand",
						Scope:              "desk",
						ActionFunctionName: "testActionFunction",
					},
				}}),
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusCreated,
			},
			beforeTest: func() map[string]string {
				res := make(map[string]string)
				res["appId"] = createRemoteApp()
				return res
			},
		},
		{
			name: "invoke function in admin",
			req: request{
				method: "PUT",
				path:   "/admin/apps/{appId}/functions",
				body: toJSONMap(app.JsonFunctionRequest{
					Method: "testFunction",
					Params: []byte("{}"),
					Context: app.ChannelContext{
						Caller: app.Caller{
							Type: "manager",
							ID:   "1",
						},
						Channel: app.Channel{
							ID: "1",
						},
					},
				}),
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusOK,
				body:       toJSONMap(successfulAppServerFunctionInvoke().expectedRequests[0].expectedResponse.body),
			},
			beforeTest: func() map[string]string {
				res := make(map[string]string)
				res["appId"] = createRemoteApp()
				installApp(res["appId"], "1")
				return res
			},
			mockServers: []mockServer{
				successfulAppServerFunctionInvoke(),
			},
		},
		{
			name: "install app in desk",
			req: request{
				method: "PUT",
				path:   "/desk/v1/channels/{channelId}/app-channels/{appId}",
				header: map[string]string{
					"x-account": "1",
				},
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusOK,
			},
			beforeTest: func() map[string]string {
				res := make(map[string]string)
				res["appId"] = createRemoteApp()
				res["channelId"] = "1"
				return res
			},
		},
		{
			name: "execute command in desk",
			req: request{
				method: "PUT",
				path:   "/desk/v1/channels/1/apps/{appId}/commands/{name}",
				header: map[string]string{
					"x-account": "1",
				},
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusOK,
				body:       toJSONMap(successfulAppServerFunctionInvoke().expectedRequests[0].expectedResponse.body["result"]),
			},
			beforeTest: func() map[string]string {
				res := make(map[string]string)
				res["appId"] = createRemoteApp()
				commandName := "deskCommand"
				res["name"] = commandName

				registerCommand(res["appId"], "testFunction", commandName, "desk")
				installApp(res["appId"], "1")
				return res
			},
			mockServers: []mockServer{
				successfulAppServerFunctionInvoke(),
			},
		},
		{
			name: "update(upsert) command",
			req: request{
				method: "POST",
				path:   "/admin/apps/{appId}/commands",
				body: toJSONMap(model.Command{
					ActionFunctionName: "newActionFunctionName",
				}),
			},
			expectedResponse: expectedResponse{
				statusCode: http.StatusCreated,
			},
			beforeTest: func() map[string]string {
				res := make(map[string]string)
				res["appId"] = createRemoteApp()
				commandName := "deskCommand"

				registerCommand(res["appId"], "testFunction", commandName, "desk")
				installApp(res["appId"], "1")
				return res
			},
			mockServers: []mockServer{
				successfulAppServerFunctionInvoke(),
			},
		},
	}

	restClient := resty.New()

	beforeAll(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeTestRes := beforeEach(test)

			r := restClient.R()
			for k, v := range beforeTestRes {
				test.req.path = strings.Replace(test.req.path, fmt.Sprintf("{%s}", k), v, -1)
			}
			r.URL = fmt.Sprintf("http://localhost:%d%s", srvPort, test.req.path)
			r.Method = test.req.method
			r.SetHeaders(test.req.header)
			r.SetBody(test.req.body)

			res, err := r.Send()
			assert.NoError(t, err)

			assert.Equalf(t, test.expectedResponse.statusCode, res.StatusCode(), "response: %v", string(res.Body()))

			if test.expectedResponse.body != nil || test.expectedResponse.assertionFunc != nil {
				resMap := make(map[string]any)
				err = json.Unmarshal(res.Body(), &resMap)
				assert.NoError(t, err)

				if test.expectedResponse.body != nil {
					assert.Equal(t, test.expectedResponse.body, resMap)
				}
				if test.expectedResponse.assertionFunc != nil {
					test.expectedResponse.assertionFunc(t, resMap)
				}
			}
		})
	}
}
