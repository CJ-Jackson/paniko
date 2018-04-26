//go:generate mockgen -write_package_comment=false -package=security -source=cookie.go -destination=cookie.mock.go
//go:generate debugflag cookie.mock.go

package security

import (
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/config"
	"github.com/gorilla/securecookie"
)

type CookieHelper interface {
	Set(context ctx.Context, cookie *http.Cookie)
	Get(context ctx.Context, name string) *http.Cookie
	Delete(context ctx.Context, name string)
}

func GetCookieHelper(context ctx.BackgroundContext) CookieHelper {
	name := "cookie-helper-e0ff25bef45a0649477d9b55231e4cc1"
	if cHelper, ok := context.Ctx(name).(CookieHelper); ok {
		return cHelper
	}

	cookieCfg := config.GetConfig(context).Cookie
	cHelper := cookieHelper{
		secureCookie: securecookie.New([]byte(cookieCfg.HashKey), []byte(cookieCfg.BlockKey)),
		errorService: common.GetErrorService(context),
	}

	context.SetCtx(name, cHelper)
	return cHelper
}

type cookieHelper struct {
	secureCookie *securecookie.SecureCookie
	errorService common.ErrorService
}

func (h cookieHelper) Set(context ctx.Context, cookie *http.Cookie) {
	var err error
	cookie.Value, err = h.secureCookie.Encode(cookie.Name, cookie.Value)
	h.errorService.CheckErrorAndPanic(err)

	http.SetCookie(context.Response(), cookie)
}

func (h cookieHelper) Get(context ctx.Context, name string) *http.Cookie {
	req := context.Request()
	cookie, err := req.Cookie(name)
	if nil != err {
		return nil
	}

	err = h.secureCookie.Decode(name, cookie.Value, &cookie.Value)
	if nil != err {
		return nil
	}

	return cookie
}

func (h cookieHelper) Delete(context ctx.Context, name string) {
	http.SetCookie(context.Response(), &http.Cookie{
		Name:   name,
		MaxAge: -1,
	})
}
