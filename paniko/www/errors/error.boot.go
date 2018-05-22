package errors

import (
	"fmt"
	"net/http"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

func bootError(controller ErrorController, muxer *httprouter.Router) {
	showError := func(req *http.Request, code int, message string) {
		controller.ShowError(ctx.GetContext(req), code, http.StatusText(code), message)
	}

	muxer.NotFound = http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
		shared.CheckIfUser(ctx.GetContext(req))
		showError(req, http.StatusNotFound, "Router could not find path")
	})
	muxer.MethodNotAllowed = http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
		shared.CheckIfUser(ctx.GetContext(req))
		showError(req, http.StatusMethodNotAllowed, "Router found path, but however method is not allowed")
	})

	muxer.PanicHandler = func(res http.ResponseWriter, req *http.Request, i interface{}) {
		switch i := i.(type) {
		case common.NoError:
			// Do nothing
		case common.HttpError:
			showError(req, i.Code, i.Message)
		case common.HttpRedirectError:
			res.Header().Set("Location", i.Location)
			res.WriteHeader(i.Code)
		default:
			showError(req, http.StatusInternalServerError, fmt.Sprint(i))
		}
	}
}
