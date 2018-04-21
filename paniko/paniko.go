package paniko

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/config"
)

func Boot() {
	context := ctx.NewBackgroundContext()

	startServer(config.GetParam(context).Address, common.GetMuxer(context))
}

func startServer(address string, handler http.Handler) {
	fmt.Println("Running Server at", address)
	log.Print(http.ListenAndServe(address, http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req, _ = ctx.NewContext(req, res)
		handler.ServeHTTP(res, req)
	})))
}
