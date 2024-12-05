package proxyapi

import (
	pvt "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/private"
	pub "github.com/channel-io/ch-app-store/internal/native/proxyapi/action/public"
	"github.com/channel-io/ch-app-store/internal/util"
)

var CHANNEL_API_ROUTES = toService(util.CHANNEL_ADMIN_API).withRule(rule{
	path("api/admin/core"): hasSubPaths{
		version("v1"): hasSubPaths{
			path("channels"):     hasFunctions{pub.GetChannel},
			path("mediums"):      hasFunctions{pvt.FindOrCreateContactAndUser},
			path("direct-chats"): hasFunctions{pvt.CreateDirectChat},
			path("groups"): hasFunctions{
				pub.GetGroup,
				pub.SearchGroups,
			},
			path("users"): hasFunctions{
				pub.GetUser,
				pvt.SearchUser,
				pvt.PatchUser,
				pvt.DeleteUser,
			},
			path("user-chats"): hasFunctions{
				pub.ManageUserChat,
				pub.GetUserChat,
				pvt.SearchUserChatsByContact,
				pvt.UpdateUserChatState,
				pvt.CreateUserChat,
			},
			path("messages"): hasFunctions{
				pub.WriteGroupMessage,
				pub.WriteGroupMessageAsManager,
				pub.WriteUserChatMessage,
				pub.WriteUserChatMessageAsManager,
				pub.WriteUserChatMessageAsUser,
				pub.WriteDirectChatMessageAsManager,
			},
			path("managers"): hasFunctions{
				pub.GetManager,
				pub.BatchGetManagers,
				pub.SearchManagers,
			},
			path("plugins"): hasFunctions{
				pvt.GetPlugin,
			},
			path("commerce"): hasFunctions{
				pvt.RegisterCommerce,
				pvt.DeregisterCommerce,
			},
		},
	},
})
