package shared

import (
	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

const (
	UserDepName  = "user-e0d6c3c9f7fd79ab13e96a6c9b7ff666"
	UserDataName = "user-c5962ba014317e92666f1cf4f6c6416c"
)

type User interface {
	Username() string
	CheckIfUser()
	CheckIfGuest()
}

func CheckIfUser(context ctx.Context) {
	context.Dep(UserDepName).(common.ContextHandler)(context)
	context.Data(UserDataName).(User).CheckIfUser()
}

func CheckIfGuest(context ctx.Context) {
	context.Dep(UserDepName).(common.ContextHandler)(context)
	context.Data(UserDataName).(User).CheckIfGuest()
}
