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

	Template := template.Must(template.New("master-Template").Parse(
		`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <title>{{ .Title }} | Paniko</title>
  </head>
  <body>
	<div class="container">
		<h1>{{ .Title }}</h1>

		{{ block "content" . }}{{ end }}
	</div>

    <script src="https://code.jquery.com/jquery-3.3.1.min.js" integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8=" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
	{{ block "js" . }}{{ end }}
  </body>
</html>`,
	))

	context.SetCtx(name, Template)
	return Template
}
