//go:generate mockgen -write_package_comment=false -package=mail -source=mailer.go -destination=mailer.mock.go
//go:generate debugflag mailer.mock.go

package mail

import (
	"fmt"
	"net/smtp"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/jpoehls/gophermail"
)

type Mailer interface {
	Dispatch()
}

func NewMailer(context ctx.BackgroundContext) Mailer {
	mailConfig := common.GetConfig(context).Mail

	return mailer{
		addr: fmt.Sprint(mailConfig.Hostname, ":", mailConfig.Port),
		auth: smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Hostname),
		message: gophermail.Message{
			From:    mailConfig.From,
			To:      mailConfig.SendTo,
			Cc:      mailConfig.SendCc,
			Subject: mailConfig.Subject,
			Body:    common.GetMessage(context),
		},
	}
}

type mailer struct {
	addr    string
	auth    smtp.Auth
	message gophermail.Message
}

func (m mailer) Dispatch() {
	gophermail.SendMail(m.addr, m.auth, &m.message)
}
