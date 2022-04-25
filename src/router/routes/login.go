package routes

import (
	"net/http"
	"social-network/src/controllers"
)

var routeLogin = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Login,
	RequiresAuthentication: false,
}
