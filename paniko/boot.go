package paniko

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/config"
	"github.com/CJ-Jackson/paniko/paniko/security"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/www"
	"github.com/CJ-Jackson/paniko/paniko/www/errors"
	"github.com/CJ-Jackson/paniko/paniko/www/login"
)

func Boot() {
	context := ctx.NewBackgroundContext()

	www.Boot(context)
	errors.Boot(context)
	login.Boot(context)

	startServer(config.GetParam(context).Address, common.GetMuxer(context), getContextBoot(context))
}

func getContextBoot(context ctx.BackgroundContext) common.ContextHandler {
	csrfHandler := security.GetCsrf(context)
	userDep := security.GetUserController(context).GetDep()

	return func(context ctx.Context) {
		context.SetDep(shared.CsrfInitName, csrfHandler)
		context.SetDep(shared.UserDepName, userDep)
	}
}

func startServer(address string, handler http.Handler, boot common.ContextHandler) {
	fmt.Println("Running Server at", address)
	log.Print(http.ListenAndServe(address, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req, context := ctx.NewContext(req, res)
		boot(context)
		handler.ServeHTTP(res, req)
	})))
}
