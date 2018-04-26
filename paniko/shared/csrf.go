package shared

import (
	"html/template"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

const (
	CsrfDataName = "csrf-6ea05b4c12ad7849b84e3604b829de5f"
	CsrfInitName = "csrf-55c7e347330248fef0096a3f99713dbd"
)

type Csrf interface {
	TokenField() template.HTML
	Token() string
}

func GetCsrf(context ctx.Context) Csrf {
	context.Dep(CsrfInitName).(common.ContextHandler)(context)
	return context.Data(CsrfDataName).(Csrf)
}
