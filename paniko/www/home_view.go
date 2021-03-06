//go:generate mockgen -write_package_comment=false -package=www -source=home_view.go -destination=home_view.mock.go
//go:generate debugflag home_view.mock.go

package www

import (
	"encoding/json"
	"html/template"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/cjtoolkit/ctx"
)

type HomeView interface {
	Index(context ctx.Context, data HomeViewIndexData)
	Json(context ctx.Context, data JsonData)
}

func NewHomeView(context ctx.BackgroundContext) HomeView {
	return homeView{
		indexTemplate: buildIndexTemplate(context),
		errorService:  common.GetErrorService(context),
	}
}

type homeView struct {
	indexTemplate *template.Template
	errorService  common.ErrorService
}

func (v homeView) Index(context ctx.Context, data HomeViewIndexData) {
	type Context struct {
		ctx.Context
		Data HomeViewIndexData
	}

	context.SetTitle("Paniko")

	err := v.indexTemplate.Execute(context.ResponseWriter(), Context{
		Context: context,
		Data:    data,
	})
	v.errorService.CheckErrorAndLog(err)
}

func (v homeView) Json(context ctx.Context, data JsonData) {
	err := json.NewEncoder(context.ResponseWriter()).Encode(data)
	v.errorService.CheckErrorAndLog(err)
}
