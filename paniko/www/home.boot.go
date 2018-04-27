package www

import (
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
	"github.com/julienschmidt/httprouter"
)

func homeBoot(homeController HomeController, muxer *httprouter.Router) {
	{
		muxer.GET(uri.Home, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfUser(context)

			homeController.Index(context)
		})
	}

	{
		muxer.PUT(uri.IAmAlive, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfUser(context)

			homeController.IAmAlive(context)
		})
	}

	{
		muxer.PUT(uri.InDanger, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfUser(context)

			homeController.InDanger(context)
		})
	}
}
