package common

import (
	"io/ioutil"
	"os"

	"github.com/cjtoolkit/ctx"
)

func GetMessage(context ctx.BackgroundContext) string {
	return context.Persist("message-071262aa97887e3b10cf6c00d7462f24", func() (interface{}, error) {
		location := os.Getenv("HOME") + "/.config/paniko/message.txt"

		errorService := GetErrorService(context)

		file, err := os.Open(location)
		errorService.CheckErrorAndPanic(err)

		b, err := ioutil.ReadAll(file)
		errorService.CheckErrorAndPanic(err)

		message := string(b)

		return message, nil
	}).(string)
}
