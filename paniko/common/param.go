package common

import (
	"os"

	"github.com/CJ-Jackson/paniko/paniko/command"
	"github.com/cjtoolkit/ctx"
)

type Param struct {
	Address    string
	Production bool
}

func GetParam(context ctx.BackgroundContext) Param {
	return context.Persist("param-7bddb00d1a070092274f0139a7cadcc9", func() (interface{}, error) {
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

		return *param, nil
	}).(Param)
}
