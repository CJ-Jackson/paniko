package login

import (
	"net/url"

	"github.com/cjtoolkit/validate/vError"
	"github.com/cjtoolkit/validate/vString"
)

type LoginValidator struct {
	usernameRule []vString.ValidationRule
	passwordRule []vString.ValidationRule
}

func NewLoginValidator() LoginValidator {
	return LoginValidator{
		usernameRule: []vString.ValidationRule{vString.Mandatory(), vString.BetweenRune(3, 50)},
		passwordRule: []vString.ValidationRule{vString.Mandatory()},
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
		Valid:       false,
	}
}

func (v LoginValidator) NewValidatedLoginForm(values url.Values) LoginForm {
	var form *LoginForm
	{
		f := v.NewLoginForm()
		form = &f
	}

	form.Attempt = true

	form.Username, form.UsernameErr = vString.Validate(values.Get("username"), v.usernameRule...)
	form.Password, form.PasswordErr = vString.Validate(values.Get("password"), v.passwordRule...)

	form.Valid = vError.CheckErr(form.UsernameErr, form.PasswordErr)

	return *form
}
