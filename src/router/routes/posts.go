package routes

import (
	"net/http"
	"social-network/src/controllers"
)

var routesPosts = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.ListPosts,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodGet,
		Function:               controllers.GetPost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.GetPostsPerUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{id}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.UnlikePost,
		RequiresAuthentication: true,
	},
}
