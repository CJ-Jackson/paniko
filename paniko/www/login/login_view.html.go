package login

import (
	"html/template"

	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/cjtoolkit/ctx"
)

const (
	loginDataName = "login-96a0b2a136ff5b03b5b7345f740d2667"
)

func buildLoginTemplate(context ctx.BackgroundContext) *template.Template {
	funcMaps := template.FuncMap{
		"form": func(context ctx.Context) LoginForm { return context.Data(loginDataName).(LoginForm) },
	}

	return template.Must(shared.CloneMasterTemplate(context).Funcs(funcMaps).Parse(loginTemplate))
}

const loginTemplate = `{{ define "content" }}{{ $form := form . }}{{ $csrf := csrf . }}{{ $formHelper := formHelper }}

{{ if $form.Attempt }}
<div class="alert alert-danger">Failed to login, try again.</div>
{{ end }}

<form method="post" novalidate>
	{{ $csrf.TokenField }}{{ $formHelper.Valid $form.Valid }}
	<input type="hidden" name="uri" value="{{ $form.Uri }}">
	<div class="form-group">
		<label for="username">Username</label>
		<input class="form-control" type="text" value="{{$form.Username}}" name="username" id="username">
		{{ $formHelper.Check $form.UsernameErr }}
	</div>
	<div class="form-group">
		<label for="password">Password</label>
		<input class="form-control" type="password" value="" name="password" id="password">
		{{ $formHelper.Check $form.PasswordErr }}
	</div>
	<button type="submit" class="btn btn-primary">Submit</button>
</form>

{{ end }}`
