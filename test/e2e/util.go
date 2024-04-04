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

	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appdevfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/approlefx"

	"github.com/channel-io/ch-app-store/config"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/gintoolfx"
	adminhandlerfx "github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/adminfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/deskfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/frontfx"
	"github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/generalfx"
	publichandlerfx "github.com/channel-io/ch-app-store/fx/corefx/apifx/httpfx/publicfx"
	"github.com/channel-io/ch-app-store/fx/corefx/configfx"
	"github.com/channel-io/ch-app-store/fx/corefx/datadogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/appfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/brieffx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/commandfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/functionfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/invokelogfx"
	"github.com/channel-io/ch-app-store/fx/corefx/internalfx/nativefx"
	"github.com/channel-io/ch-app-store/fx/corefx/logfx"
	"github.com/channel-io/ch-app-store/fx/corefx/restyfx"
	"github.com/channel-io/ch-app-store/test/mockauth"
)

var httpModule = fx.Options(
	generalfx.GeneralHandlers,
	frontfx.FrontHandlers,
	deskfx.DeskHandlers,
	publichandlerfx.PublicHandlers,
	adminhandlerfx.AdminHandlers,
	gintoolfx.ApiServer,
)

var fullAppModule = fx.Options(
	datadogfx.Datadog,
	configfx.Values,
	httpModule,
	approlefx.AppRole,
	functionfx.Function,
	brieffx.Brief,
	appfx.App,
	appdevfx.AppDev,
	invokelogfx.Loggers,
	logfx.Logger,
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
			r, err := db.Query(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", tableName))
			if err != nil {
				panic(err)
			}
			defer r.Close()
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
