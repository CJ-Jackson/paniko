package common

import (
	"encoding/json"
	"os"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common/config"
)

func GetConfig(context ctx.BackgroundContext) config.Base {
	const name = "config-5a43ef4f8f6dbc0ee0ec3471d26dfdcd"
	if configVar, ok := context.Ctx(name).(config.Base); ok {
		return configVar
	}

	configVar := config.Base{
		DaysToExpiry: 7,
		CsrfKey:      "",
		Mail: config.Mail{
			SendTo:  []string{},
			SendCc:  []string{},
			Subject: "",
		},
		Password: config.Password{
			Location: os.Getenv("HOME") + "/.config/paniko/password.json",
		},
		Cookie: config.Cookie{
			HashKey:  "",
			BlockKey: "",
		},
	}
	location := os.Getenv("HOME") + "/.config/paniko/config.json"

	errorService := GetErrorService(context)

	file, err := os.Open(location)
	errorService.CheckErrorAndPanic(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&configVar)
	errorService.CheckErrorAndPanic(err)

	context.SetCtx(name, configVar)
	return configVar
}
