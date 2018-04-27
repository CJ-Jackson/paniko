package login

import (
	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/security"
)

type LoginController struct {
	userController security.UserController
	loginView      LoginView
}

func NewLoginController(context ctx.BackgroundContext) LoginController {
	return LoginController{
		userController: security.GetUserController(context),
		loginView:      NewLoginView(context),
	}
}

func (c LoginController) ShowLogin(context ctx.Context, form LoginForm) {
	c.loginView.LoginTemplate(context, form)
}

func (c LoginController) DoLogin(context ctx.Context, form LoginForm) {
	if c.userController.Login(context, form.Username, form.Password) {
		common.HaltSeeOther(form.Uri)
	}

	c.ShowLogin(context, form)
}
