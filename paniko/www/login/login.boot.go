package login

import (
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
	"github.com/julienschmidt/httprouter"
)

func bootLogin(controller LoginController, muxer *httprouter.Router) {
	// Log In
	{
		muxer.GET(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfGuest(context)
			shared.InitCsrf(context)

			controller.ShowLogin(context, LoginForm{
				Uri:      "/",
				Username: "",
				Password: "",
				Attempt:  false,
			})
		})
		muxer.POST(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.CheckIfGuest(context)
			request = context.Request()
			request.ParseForm()
			shared.CheckCsrf(context)

			controller.DoLogin(context, LoginForm{
				Uri:      "/",
				Username: request.PostForm.Get("username"),
				Password: request.PostForm.Get("password"),
				Attempt:  true,
			})
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
