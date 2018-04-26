package config

import (
	"encoding/json"
	"os"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

type Config struct {
	DaysToExpiry int
	CsrfKey      string
	Mail         Mail
	Password     Password
}

type Mail struct {
	SendTo  []string
	SendCc  []string
	Subject string
}

type Password struct {
	Salt     string
	Location string
}

func GetConfig(context ctx.BackgroundContext) Config {
	name := "config-5a43ef4f8f6dbc0ee0ec3471d26dfdcd"
	if config, ok := context.Ctx(name).(Config); ok {
		return config
	}

	config := Config{
		Password: Password{
			Location: os.Getenv("HOME") + "/.config/paniko/password.json",
		},
	}
	location := os.Getenv("HOME") + "/.config/paniko/config.json"

	file, err := os.Open(location)
	common.CheckErrorAndExit(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	common.CheckErrorAndExit(err)

	context.SetCtx(name, config)
	return config
}
