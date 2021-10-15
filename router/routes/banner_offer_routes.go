package routes

import (
	"astro/controllers"
	"net/http"
)

var BannerOfferRoutes = []Route{
	{
		Uri:          "/banneroffer",
		Method:       http.MethodPost,
		Handler:      controllers.CreateBannerOffer,
	},
	{
		Uri:          "/banneroffer",
		Method:       http.MethodGet,
		Handler:      controllers.GetBannerOffer,
	},
}
