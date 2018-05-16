package security

import (
	"html/template"
	"net/http"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/cjtoolkit/ctx"
	"github.com/gorilla/csrf"
)

func GetCsrf(context ctx.BackgroundContext) common.ContextHandler {
	const name = "csrf-7cf4ddd5429f9237ebd331a9c65498ed"
	if contextHandler, ok := context.Get(name).(common.ContextHandler); ok {
		return contextHandler
	}

	csrfProtect := csrf.Protect(
		[]byte(common.GetConfig(context).CsrfKey),
		csrf.Secure(false),
		csrf.ErrorHandler(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
			common.HaltForbidden("Invalid Csrf Token")
		})),
	)

	contextHandler := func(context ctx.Context) {
		if _, ok := context.Data(shared.CsrfDataName).(csrfData); ok {
			return
		}

		csrfProtect(http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
			context.SetData(shared.CsrfDataName, shared.Csrf(csrfData{
				tokenField: csrf.TemplateField(req),
				token:      csrf.Token(req),
			}))
		})).ServeHTTP(context.ResponseWriter(), context.Request())
	}

	context.Set(name, contextHandler)
	return contextHandler
}

type csrfData struct {
	tokenField template.HTML
	token      string
}

func (d csrfData) TokenField() template.HTML { return d.tokenField }

func (d csrfData) Token() string { return d.token }
