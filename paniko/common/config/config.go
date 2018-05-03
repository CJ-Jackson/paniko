package config

import "net/mail"

type Base struct {
	DaysToExpiry int
	CsrfKey      string
	Mail         Mail
	Password     Password
	Cookie       Cookie
}

type Mail struct {
	Username string
	Password string
	Hostname string
	Port     int
	From     mail.Address
	SendTo   []mail.Address
	SendCc   []mail.Address
	Subject  string
}

type Password struct {
	Salt     string
	Location string
}

type Cookie struct {
	HashKey  string
	BlockKey string
}
