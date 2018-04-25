package config

import (
	"os"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/command"
)

type Param struct {
	Address    string
	Production bool
}

func GetParam(context ctx.BackgroundContext) Param {
	name := "param-7bddb00d1a070092274f0139a7cadcc9"
	if param, ok := context.Ctx(name).(Param); ok {
		return param
	}

	param := &Param{
		Address:    ":8080",
		Production: false,
	}

	options := command.BuildOptions(os.Args[1:])

	options.ExecOption("address", func(strings []string) {
		param.Address = strings[0]
	})
	options.ExecOption("prod", func(_ []string) {
		param.Production = true
	})

	context.SetCtx(name, *param)
	return *param
}
