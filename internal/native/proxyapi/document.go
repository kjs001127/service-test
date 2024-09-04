package proxyapi

import (
	"github.com/channel-io/ch-app-store/internal/native/proxyapi/action/public"
	"github.com/channel-io/ch-app-store/internal/util"
)

var DOCUMENT_API_ROUTES = toService(util.DOCUMENT_API).withRule(rule{
	path("internal-rpc"): hasSubPaths{
		version("v1"): hasFunctions{
			public.GetArticle,
			public.GetRevision,
			public.SearchArticles,
		},
	},
})
