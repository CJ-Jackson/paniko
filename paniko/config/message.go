package config

import (
	"io/ioutil"
	"os"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

func GetMessage(context ctx.BackgroundContext) string {
	name := "message-071262aa97887e3b10cf6c00d7462f24"
	if message, ok := context.Ctx(name).(string); ok {
		return message
	}

	location := os.Getenv("HOME") + "/.config/paniko/message.txt"

	errorService := common.GetErrorService(context)

	file, err := os.Open(location)
	errorService.CheckErrorAndPanic(err)

	b, err := ioutil.ReadAll(file)
	errorService.CheckErrorAndPanic(err)

	message := string(b)

	context.SetCtx(name, message)
	return message
}
