package errors

import (
	"html/template"
)

type ErrorTemplateData struct {
	Production bool
	StackTrace string
	Message    string
}

func buildErrorTemplate() *template.Template {
	return template.Must(template.New("error").Parse(errorTemplate))
}

const errorTemplate = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">

    <title>{{ .Title }} | Paniko</title>
  </head>
  <body>
		<h1>{{ .Title }}</h1>

		{{ if not .Data.Production }}
		<p>{{ .Data.Message }}</p>

		<pre>{{ .Data.StackTrace }}</pre>
		{{ end }}
  </body>
</html>`
