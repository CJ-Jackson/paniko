package common

import (
	"io/ioutil"
	"os"

	"github.com/CJ-Jackson/ctx"
)

func GetMessage(context ctx.BackgroundContext) string {
	const name = "message-071262aa97887e3b10cf6c00d7462f24"
	if message, ok := context.Ctx(name).(string); ok {
		return message
	}

	location := os.Getenv("HOME") + "/.config/paniko/message.txt"

	errorService := GetErrorService(context)

	file, err := os.Open(location)
	errorService.CheckErrorAndPanic(err)

	b, err := ioutil.ReadAll(file)
	errorService.CheckErrorAndPanic(err)

	message := string(b)

	context.SetCtx(name, message)
	return message
}
