package routes

import (
	"net/http"
	"social-network/src/controllers"
)

var routesUsers = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Request:                controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Request:                controllers.ListUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodGet,
		Request:                controllers.GetUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodPut,
		Request:                controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{id}",
		Method:                 http.MethodDelete,
		Request:                controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
