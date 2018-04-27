//go:generate mockgen -write_package_comment=false -package=security -source=user.go -destination=user.mock.go
//go:generate debugflag user.mock.go

package security

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/config"
	"github.com/CJ-Jackson/paniko/paniko/shared"
)

const userCookieName = "user"

type UserController interface {
	CheckCookie(context ctx.Context) shared.User
	GetDep() common.ContextHandler
	Login(context ctx.Context, username, password string) bool
	UpdateUser(username, password string)
	Save()
	Logout(context ctx.Context)
}

type userController struct {
	passwordLocation string
	password         common.Password
	users            map[string]string
	errorService     common.ErrorService
	cookieHelper     CookieHelper
}

func GetUserController(context ctx.BackgroundContext) userController {
	name := "user-cbb169f923ea34417740c8d0cda6bc16"

	passwordCfg := config.GetConfig(context).Password
	errorService := common.GetErrorService(context)
	password := common.NewPassword(passwordCfg.Salt, errorService)

	users := map[string]string{
		"admin": password.SaltPassword("password"),
	}

	file, err := os.Open(passwordCfg.Location)
	if nil != err {
		file, err := os.Create(passwordCfg.Location)
		defer file.Close()
		errorService.CheckErrorAndPanic(err)

		err = json.NewEncoder(file).Encode(users)
		errorService.CheckErrorAndPanic(err)
	} else {
		defer file.Close()

		err = json.NewDecoder(file).Decode(&users)
		errorService.CheckErrorAndPanic(err)
	}

	userC := userController{
		passwordLocation: passwordCfg.Location,
		password:         password,
		users:            users,
		errorService:     errorService,
		cookieHelper:     GetCookieHelper(context),
	}

	context.SetCtx(name, userC)
	return userC
}

func (c userController) CheckCookie(context ctx.Context) shared.User {
	cookie := c.cookieHelper.Get(context, userCookieName)
	if nil == cookie {
		return Guest{}
	}

	values := strings.Split(cookie.Value, ":")
	if len(values) < 2 {
		return Guest{}
	}
	username := values[0]
	password := strings.Join(values[1:], ":")

	if password != c.users[username] {
		return Guest{}
	}

	return User{
		username: username,
	}
}

func (c userController) GetDep() common.ContextHandler {
	return func(context ctx.Context) {
		if _, ok := context.Data(shared.UserDataName).(shared.User); ok {
			return
		}

		context.SetData(shared.UserDataName, c.CheckCookie(context))
	}
}

func (c userController) Login(context ctx.Context, username, password string) bool {
	if c.password.CheckPassword(password, c.users[username]) {
		c.cookieHelper.Set(context, &http.Cookie{
			Name:    userCookieName,
			Expires: time.Now().AddDate(0, 1, 0),
			Value:   username + ":" + c.password.SaltPassword(password),
		})
		return true
	}

	return false
}

func (c userController) UpdateUser(username, password string) {
	c.users[username] = c.password.SaltPassword(password)
}

func (c userController) Save() {
	file, err := os.Create(c.passwordLocation)
	defer file.Close()
	c.errorService.CheckErrorAndPanic(err)

	err = json.NewEncoder(file).Encode(c.users)
	c.errorService.CheckErrorAndPanic(err)
}

func (c userController) Logout(context ctx.Context) {
	c.cookieHelper.Delete(context, userCookieName)
}
