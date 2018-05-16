package shared

import (
	"html/template"

	"github.com/cjtoolkit/ctx"
)

func CloneMasterTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(getMasterTemplate(context).Clone())
}

func getMasterTemplate(context ctx.BackgroundContext) *template.Template {
	return context.Persist("master-template-73e8c809b3a0930c4d0085f5d183a6ab", func() (interface{}, error) {
		funcMaps := template.FuncMap{
			"csrf": GetCsrf,
			"user": GetUser,
		}

		Template := template.Must(template.New("master-Template").Funcs(funcMaps).Parse(masterTemplate))

		return Template, nil
	}).(*template.Template)
}
