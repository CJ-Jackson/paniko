package login

import (
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
	"github.com/julienschmidt/httprouter"
)

func bootLogin(controller LoginController, muxer *httprouter.Router) {
	{
		muxer.GET(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			shared.InitCsrf(context)

			controller.ShowLogin(context, LoginForm{
				Uri:      request.Header.Get("Referer"),
				Username: "",
				Password: "",
				Attempt:  false,
			})
		})
		muxer.POST(uri.Login, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			context := ctx.GetContext(request)
			request = context.Request()
			request.ParseForm()
			shared.CheckCsrf(context)

			controller.DoLogin(context, LoginForm{
				Uri:      request.PostForm.Get("uri"),
				Username: request.PostForm.Get("username"),
				Password: request.PostForm.Get("password"),
				Attempt:  true,
			})
		})
	}
}
