package e2e

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	"github.com/channel-io/ch-app-store/internal/command/model"
)

var cli = resty.New()

const baseUrl = "http://localhost:3020"
const FakeFunctionUrl = "http://www.function.com"
const FakeWamUrl = "http://www.wam.com"

func createRemoteApp() string {
	req := make(map[string]any)
	req["isPrivate"] = false
	req["configSchemas"] = []string{}
	req["roles"] = []string{}
	req["wamUrl"] = FakeWamUrl
	req["functionUrl"] = FakeFunctionUrl

	requester := cli.R()
	requester.SetBody(req)
	res, err := requester.Post(baseUrl + "/admin/apps")
	if err != nil {
		panic(err)
	}

	var resMap map[string]any
	err = json.Unmarshal(res.Body(), &resMap)
	if err != nil {
		panic(err)
	}
	return resMap["id"].(string)
}

func registerCommand(appID, functionName, name string, scope model.Scope) {
	req := dto.RegisterRequest{Commands: []*model.Command{
		{
			AppID:              appID,
			Name:               name,
			Scope:              scope,
			AlfMode:            "NOT_IMPORTANT",
			ActionFunctionName: functionName,
		},
	}}

	requester := cli.R()
	requester.SetBody(req)
	res, err := requester.Post(baseUrl + "/admin/apps/" + appID + "/commands")
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != http.StatusCreated {
		panic("response code is not 201")
	}
}

func installApp(appID, channelID string) {
	requester := cli.R()
	requester.SetHeader("x-account", "1")
	res, err := requester.Put(baseUrl + "/desk/v1/channels/" + channelID + "/installed-apps/" + appID)
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != http.StatusOK {
		panic("response code is not 200")
	}
}
