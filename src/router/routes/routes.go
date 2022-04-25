package routes

import "net/http"

type Route struct {
	URI                    string
	Method                 string
	Request                func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
