package shared

import (
	"html/template"

	"github.com/CJ-Jackson/ctx"
)

func CloneMasterTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(getMasterTemplate(context).Clone())
}

func getMasterTemplate(context ctx.BackgroundContext) *template.Template {
	name := "master-template-73e8c809b3a0930c4d0085f5d183a6ab"
	if Template, ok := context.Ctx(name).(*template.Template); ok {
		return Template
	}

	funcMaps := template.FuncMap{
		"csrf": GetCsrf,
	}

	Template := template.Must(template.New("master-Template").Funcs(funcMaps).Parse(masterTemplate))

	context.SetCtx(name, Template)
	return Template
}
