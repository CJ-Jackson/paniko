package login

import (
	"net/url"

	"github.com/cjtoolkit/validate/vError"
	"github.com/cjtoolkit/validate/vString"
)

type LoginValidator struct {
	usernameRules []vString.ValidationRule
	passwordRules []vString.ValidationRule
}

func NewLoginValidator() LoginValidator {
	return LoginValidator{
		usernameRules: []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(3, 50)},
		passwordRules: []vString.ValidationRule{vString.Mandatory()},
	}
}

func (v LoginValidator) NewLoginForm() LoginForm {
	return LoginForm{
		Uri:         "",
		Username:    "",
		UsernameErr: nil,
		Password:    "",
		PasswordErr: nil,
		Attempt:     false,
		Valid:       true,
	}
}

func (v LoginValidator) NewValidatedLoginForm(values url.Values) LoginForm {
	var form *LoginForm
	{
		f := v.NewLoginForm()
		form = &f
	}

	form.Attempt = true

	form.Username, form.UsernameErr = vString.Validate(values.Get("username"), v.usernameRules...)
	form.Password, form.PasswordErr = vString.Validate(values.Get("password"), v.passwordRules...)

	form.Valid = vError.CheckErr(form.UsernameErr, form.PasswordErr)

	return *form
}
