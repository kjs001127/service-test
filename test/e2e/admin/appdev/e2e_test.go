package appdev_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/channel-io/ch-app-store/api/http/admin/dto"
	mockaccount "github.com/channel-io/ch-app-store/generated/mock/auth/principal/account"
	app "github.com/channel-io/ch-app-store/internal/app/model"
	appmodel "github.com/channel-io/ch-app-store/internal/app/model"
	"github.com/channel-io/ch-app-store/internal/appdev/svc"
	functionmodel "github.com/channel-io/ch-app-store/internal/apphttp/model"
	"github.com/channel-io/ch-app-store/internal/auth/principal/account"
	"github.com/channel-io/ch-app-store/internal/command/model"
	. "github.com/channel-io/ch-app-store/test/e2e"
	. "github.com/channel-io/ch-app-store/test/e2e/admin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	baseURL          = "http://localhost:3020"
	fetchManagerPath = "http://localhost:8080/api/admin/manager"
	channelID        = "1"
	fakeXAccount     = "x-account"
	managerID        = "1253"
)

type E2ETest struct {
	suite.Suite

	test           *TestApp
	managerFetcher *mockaccount.ManagerFetcher
}

func (e *E2ETest) SetupSuite() {
	e.managerFetcher = mockaccount.NewManagerFetcher(e.T())
	e.test = NewTestApp(
		Mock[account.ManagerFetcher](e.managerFetcher),
	)
}

func (e *E2ETest) TearDownSuite() {
	fmt.Println("TearDownTest")
	e.test.Stop()
	e.test.TruncateAll()
}

func (e *E2ETest) TestPing() {
	m := NewHttpTestClient(baseURL)
	response, _ := m.Get("/ping").Do()

	assert.Equal(e.T(), http.StatusOK, response.StatusCode())
	assert.Equal(e.T(), "pong", string(response.Body()))
}

func (e *E2ETest) TestCreateRemoteApp() app.App {
	fakeFunctionUrl := "http://www.function.com"
	fakeWamUrl := "http://www.wam.com"

	m := NewHttpTestClient(baseURL)

	appCreateRequest := appmodel.App{
		Title:     "TestApp",
		IsPrivate: false,
	}

	urlCreateRequest := functionmodel.Urls{
		FunctionURL: &fakeFunctionUrl,
		WamURL:      &fakeWamUrl,
	}

	body := svc.AppRequest{
		Roles: nil,
		RemoteApp: &svc.RemoteApp{
			App:  &appCreateRequest,
			Urls: urlCreateRequest,
		},
	}

	raw, _ := json.Marshal(body)

	response, _ := m.Post("/admin/apps").SetBody(raw).Do()

	var appResponse svc.AppResponse
	_ = json.Unmarshal(response.Body(), &appResponse)

	assert.Equal(e.T(), http.StatusCreated, response.StatusCode())
	assert.Equal(e.T(), "TestApp", appResponse.Title)
	assert.Equal(e.T(), false, appResponse.IsPrivate)
	assert.Equal(e.T(), *appResponse.FunctionURL, fakeFunctionUrl)
	return *appResponse.App
}

func (e *E2ETest) TestRegisterCommand() {

	m := NewHttpTestClient(baseURL)
	createdApp := e.TestCreateRemoteApp()

	commandBody := dto.RegisterRequest{
		Commands: []*model.Command{
			{
				Name:               "testCommand",
				Scope:              "desk",
				ActionFunctionName: "testActionFunction",
				AppID:              createdApp.ID,
			},
		},
		EnableByDefault: true,
	}
	raw, _ := json.Marshal(commandBody)
	response, _ := m.Post("/admin/apps/" + createdApp.ID + "/commands").SetBody(raw).Do()

	assert.Equal(e.T(), http.StatusCreated, response.StatusCode())
}

func (e *E2ETest) TestInstallApp() {
	m := NewHttpTestClient(baseURL)
	createdApp := e.TestCreateRemoteApp()

	managerPrincipal := account.ManagerPrincipal{
		Manager: account.Manager{
			ID: managerID,
		},
		Token: fakeXAccount,
	}

	e.managerFetcher.EXPECT().FetchManager(mock.Anything, channelID, fakeXAccount).Return(managerPrincipal, nil)

	res, err := m.Put("/desk/v1/channels/"+channelID+"/installed-apps/"+createdApp.ID).
		SetHeader("x-account", fakeXAccount).
		Do()

	assert.Equal(e.T(), http.StatusOK, res.StatusCode())
	assert.Nil(e.T(), err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(E2ETest))
}
