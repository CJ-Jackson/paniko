package www

import (
	"html/template"

	"encoding/json"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
)

type HomeView interface {
	Index(context ctx.Context, data HomeViewIndexData)
	Json(context ctx.Context, data JsonData)
}

func NewHomeView(context ctx.BackgroundContext) HomeView {
	indexName := "index-08a83f56f55ae2aed3d2b78c58d2645c"
	indexTemplate := template.Must(shared.CloneMasterTemplate(context).Funcs(template.FuncMap{
		"index": func(context ctx.Context) HomeViewIndexData {
			return context.Data(indexName).(HomeViewIndexData)
		},
	}).Parse(`
{{define "content"}}
Hello
{{end}}
`))

	return homeView{
		indexName:     indexName,
		indexTemplate: indexTemplate,
	}
}

type homeView struct {
	indexName     string
	indexTemplate *template.Template
}

type HomeViewIndexData struct {
	Expiry string
}

func (v homeView) Index(context ctx.Context, data HomeViewIndexData) {
	context.SetTitle("Home")
	context.SetData(v.indexName, data)

	v.indexTemplate.Execute(context.Response(), context)
}

type JsonData struct {
	Alert string
	When  string
}

func (v homeView) Json(context ctx.Context, data JsonData) {
	json.NewEncoder(context.Response()).Encode(data)
}
