package e2e

import (
	"github.com/channel-io/ch-app-store/internal/command/svc"
)

func successfulAppServerFunctionInvoke() mockServer {
	res := &svc.Action{
		Type:       svc.ActionType("button"),
		Attributes: map[string]any{"fakeAttribute": "fakeValue"},
	}
	return mockServer{
		url: FakeFunctionUrl,
		expectedRequests: []struct {
			req              request
			expectedResponse expectedResponse
		}{
			{
				req: request{
					method: "PUT",
					path:   "",
					header: map[string]string{"Content-Type": "application/json"},
				},
				expectedResponse: expectedResponse{
					statusCode: 200,
					body:       map[string]any{"result": res},
				},
			},
		},
	}
}
