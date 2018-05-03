package errors

import (
	"runtime/debug"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

type ErrorController struct {
	production bool
	view       ErrorView
}

func NewErrorController(context ctx.BackgroundContext) ErrorController {
	return ErrorController{
		production: common.GetParam(context).Production,
		view:       NewErrorView(context),
	}
}

func (c ErrorController) ShowError(context ctx.Context, code int, status, message string) {
	stackTrace := []byte{}
	if !c.production {
		stackTrace = debug.Stack()
	}

	c.view.ErrorTemplate(context, code, status, ErrorTemplateData{
		Production: c.production,
		StackTrace: string(stackTrace),
		Message:    message,
	})
}
