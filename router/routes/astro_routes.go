package routes

import (
	"astro/controllers"
	"net/http"
)

var AstroRoutes = []Route{
	{
		Uri:          "/astro",
		Method:       http.MethodPost,
		Handler:      controllers.CreateAstro,
	},
	{
		Uri:          "/astro",
		Method:       http.MethodGet,
		Handler:      controllers.GetAstro,
	},
}
