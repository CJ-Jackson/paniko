package login

import (
	"net/http"

	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

func bootLogin(controller LoginController, validator LoginValidator, muxer *httprouter.Router) {
	// Log In
	{
		muxer.GET(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfGuest(context)
			shared.InitCsrf(context)

			controller.ShowLogin(context, validator.NewLoginForm())
		})
		muxer.POST(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfGuest(context)
			request.ParseForm()
			shared.CheckCsrf(context)

			controller.DoLogin(context, validator.NewValidatedLoginForm(request.PostForm))
		})
	}

	// Log out
	{
		muxer.GET(uri.Logout, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfUser(context)

			controller.DoLogout(context)
		})
	}
}
