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

func GetUser(context ctx.Context) User {
	context.Dep(UserDepName).(common.ContextHandler)(context)
	return context.Data(UserDataName).(User)
}

func CheckIfUser(context ctx.Context) {
	GetUser(context).CheckIfUser()
}

func CheckIfGuest(context ctx.Context) {
	GetUser(context).CheckIfGuest()
}
