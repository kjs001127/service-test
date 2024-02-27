package e2e

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/channel-io/go-lib/pkg/log"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"go.uber.org/fx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/gintoolfx"
	adminhandlerfx "github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/generalfx"
	publichandlerfx "github.com/channel-io/ch-app-store/fx/commonfx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/configfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/dbfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/internalfx/remoteappfx"
	"github.com/channel-io/ch-app-store/fx/commonfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/commonfx/restyfx"
	"github.com/channel-io/ch-app-store/test/mockauth"
)

var httpModule = fx.Module(
	"httpModule",
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publichandlerfx.PublicHandlers,
	adminhandlerfx.AdminHandlers,
	gintoolfx.ApiServer,
)

var fullAppModule = fx.Module(
	"appTest",
	dbfx.Postgres,
	configfx.Values,
	httpModule,
	remoteappfx.RemoteAppDev,
	datadogfx.Datadog,
	brieffx.Brief,
	appfx.App,
	mockauth.Module,
	commandfx.Command,
	restyfx.Clients,
	nativefx.Native,

	fx.Supply(log.New("Test")),
)

func truncateDB(db *sql.DB) {
	rows, err := db.Query(fmt.Sprintf("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname='%s'", config.Get().Psql.Schema))
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		if tableName != "flyway_schema_history" {
			db.Query(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", tableName))
		}
	}
}

var internalClients []*resty.Client
var db *sql.DB

func beforeAll(t *testing.T) {
	var options []fx.Option
	initDB := fx.Invoke(truncateDB)
	options = append(options, fullAppModule, initDB, fx.Supply(t))
	mockServers := fx.Invoke(
		fx.Annotate(func(dwCli *resty.Client, appCli *resty.Client) {
			httpmock.ActivateNonDefault(dwCli.GetClient())
			httpmock.ActivateNonDefault(appCli.GetClient())
			internalClients = append(internalClients, dwCli, appCli)
		}, fx.ParamTags(restyfx.Dw, restyfx.App)),
	)
	getDB := fx.Invoke(func(testDB *sql.DB) { db = testDB })
	options = append(options, mockServers, getDB)

	fx.New(options...)
}

func beforeEach(testInfo testInfo) map[string]string {
	truncateDB(db)
	for _, srv := range testInfo.mockServers {
		for _, req := range srv.expectedRequests {
			responder, err := httpmock.NewJsonResponder(req.expectedResponse.statusCode, req.expectedResponse.body)
			if err != nil {
				panic(err)
			}
			httpmock.RegisterResponder(req.req.method, srv.url+req.req.path, responder)
		}
	}

	if testInfo.beforeTest != nil {
		return testInfo.beforeTest()
	}
	return nil
}

func toJSONMap(input any) map[string]any {
	var output map[string]any
	buf, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	_ = json.Unmarshal(buf, &output)
	return output
}