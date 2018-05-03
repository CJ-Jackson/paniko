package config

type Config struct {
	DaysToExpiry int
	CsrfKey      string
	Mail         Mail
	Password     Password
	Cookie       Cookie
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

type Cookie struct {
	HashKey  string
	BlockKey string
}
