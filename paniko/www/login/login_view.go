//go:generate mockgen -write_package_comment=false -package=login -source=login_view.go -destination=login_view.mock.go
//go:generate debugflag login_view.mock.go

package login

import (
	"html/template"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

type LoginView interface {
	LoginTemplate(context ctx.Context, form LoginForm)
}

func NewLoginView(context ctx.BackgroundContext) LoginView {
	return loginView{
		errorService:  common.GetErrorService(context),
		loginTemplate: buildLoginTemplate(context),
	}
}

type loginView struct {
	errorService  common.ErrorService
	loginTemplate *template.Template
}

func (v loginView) LoginTemplate(context ctx.Context, form LoginForm) {
	context.SetTitle("Login")
	context.SetData(loginDataName, form)

	err := v.loginTemplate.Execute(context.Response(), context)
	v.errorService.CheckErrorAndLog(err)
}
