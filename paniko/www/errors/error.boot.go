package errors

import (
	"fmt"
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/julienschmidt/httprouter"
)

func bootError(controller ErrorController, muxer *httprouter.Router) {
	showError := func(req *http.Request, code int, message string) {
		controller.ShowError(ctx.GetContext(req), code, http.StatusText(code), message)
	}

	muxer.NotFound = func(_ http.ResponseWriter, req *http.Request) {
		showError(req, http.StatusNotFound, "Router could not find path")
	}
	muxer.MethodNotAllowed = func(_ http.ResponseWriter, req *http.Request) {
		showError(req, http.StatusMethodNotAllowed, "Router found path, but however method is not allowed")
	}

	muxer.PanicHandler = func(_ http.ResponseWriter, req *http.Request, i interface{}) {
		switch i := i.(type) {
		case common.NoError:
			// Do nothing
		case common.HttpError:
			showError(req, i.Code, i.Message)
		default:
			showError(req, http.StatusInternalServerError, fmt.Sprint(i))
		}
	}
}
