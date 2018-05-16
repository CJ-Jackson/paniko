//go:generate mockgen -write_package_comment=false -package=errors -source=error_view.go -destination=error_view.mock.go
//go:generate debugflag error_view.mock.go

package errors

import (
	"html/template"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/cjtoolkit/ctx"
)

type ErrorView interface {
	ErrorTemplate(context ctx.Context, code int, title string, data ErrorTemplateData)
}

type errorView struct {
	errorService  common.ErrorService
	errorTemplate *template.Template
}

func NewErrorView(context ctx.BackgroundContext) ErrorView {
	return errorView{
		errorService:  common.GetErrorService(context),
		errorTemplate: buildErrorTemplate(),
	}
}

func (v errorView) ErrorTemplate(context ctx.Context, code int, title string, data ErrorTemplateData) {
	context.SetTitle(title)
	context.SetData(errorTemplateDataName, data)

	res := context.ResponseWriter()
	res.WriteHeader(code)

	err := v.errorTemplate.Execute(res, context)
	v.errorService.CheckErrorAndLog(err)
}
