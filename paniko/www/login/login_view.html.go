package login

import (
	"html/template"

	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/cjtoolkit/ctx"
)

func buildLoginTemplate(context ctx.BackgroundContext) *template.Template {
	return template.Must(shared.CloneMasterTemplate(context).Parse(loginTemplate))
}

const loginTemplate = `{{ define "content" }}{{ $csrf := csrf . }}{{ $fH := formHelper }}

{{ if .Form.Attempt }}
<div class="alert alert-danger">Failed to login, try again.</div>
{{ end }}

<form method="post" novalidate>
	{{ $csrf.TokenField }}{{ $fH.Set .Form.Attempt .Form.Valid }}
	<input type="hidden" name="uri" value="{{ .Form.Uri }}">
	<div class="form-group">
		<label for="username">Username</label>
		<input class="form-control {{ $fH.ValidClass .Form.UsernameErr }}" type="text" value="{{.Form.Username}}" name="username" id="username">
		{{ $fH.CheckErr .Form.UsernameErr }}
	</div>
	<div class="form-group">
		<label for="password">Password</label>
		<input class="form-control {{ $fH.ValidClass .Form.PasswordErr }}" type="password" value="" name="password" id="password">
		{{ $fH.CheckErr .Form.PasswordErr }}
	</div>
	<button type="submit" class="btn btn-primary">Submit</button>
</form>

{{ end }}`
