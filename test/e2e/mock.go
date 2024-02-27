package e2e

import (
	"github.com/channel-io/ch-app-store/internal/command/domain"
)

func successfulAppServerFunctionInvoke() mockServer {
	res := &domain.Action{
		Type:       domain.ActionType("button"),
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
