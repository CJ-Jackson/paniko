package errors

import (
	"html/template"

	"github.com/CJ-Jackson/ctx"
)

const errorTemplateDataName = "error-e4e93993b5d574d784ac512857d266c4"

type ErrorTemplateData struct {
	Production bool
	StackTrace string
	Message    string
}

func buildErrorTemplate() *template.Template {
	funcMaps := template.FuncMap{
		"error": func(context ctx.Context) ErrorTemplateData {
			return context.Data(errorTemplateDataName).(ErrorTemplateData)
		},
	}

	return template.Must(template.New("error").Funcs(funcMaps).Parse(errorTemplate))
}

const errorTemplate = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">

    <title>{{ .Title }} | Paniko</title>
  </head>
  <body>
		<h1>{{ .Title }}</h1>

		{{ $error := error . }}
		{{ if not $error.Production }}
		<p>{{ $error.Message }}</p>

		<pre>{{ $error.StackTrace }}</pre>
		{{ end }}
  </body>
</html>`
