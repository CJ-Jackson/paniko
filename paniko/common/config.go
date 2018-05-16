package common

import (
	"encoding/json"
	"net/mail"
	"os"

	"github.com/CJ-Jackson/paniko/paniko/common/config"
	"github.com/cjtoolkit/ctx"
)

func GetConfig(context ctx.BackgroundContext) config.Base {
	return context.Persist("config-5a43ef4f8f6dbc0ee0ec3471d26dfdcd", func() (interface{}, error) {
		configVar := config.Base{
			DaysToExpiry: 7,
			CsrfKey:      "",
			Mail: config.Mail{
				Username: "",
				Password: "",
				Hostname: "",
				Port:     25,
				From:     mail.Address{},
				SendTo:   []mail.Address{},
				SendCc:   []mail.Address{},
				Subject:  "",
			},
			Password: config.Password{
				Salt:     "",
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

		return configVar, nil
	}).(config.Base)
}
