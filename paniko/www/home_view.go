package www

import (
	"html/template"

	"encoding/json"

	"github.com/CJ-Jackson/ctx"
)

type HomeView interface {
	Index(context ctx.Context, data HomeViewIndexData)
	Json(context ctx.Context, data JsonData)
}

func NewHomeView(context ctx.BackgroundContext) HomeView {
	return homeView{
		indexTemplate: buildIndexTemplate(context),
	}
}

type homeView struct {
	indexTemplate *template.Template
}

func (v homeView) Index(context ctx.Context, data HomeViewIndexData) {
	context.SetTitle("Paniko")
	context.SetData(indexName, data)

	v.indexTemplate.Execute(context.Response(), context)
}

func (v homeView) Json(context ctx.Context, data JsonData) {
	json.NewEncoder(context.Response()).Encode(data)
}
