package mail

import (
	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/expiration"
)

type Dispatcher struct {
	mailer     Mailer
	expiration expiration.Expiration
}

func NewDispatcher(context ctx.BackgroundContext) Dispatcher {
	return Dispatcher{
		mailer:     NewMailer(context),
		expiration: expiration.GetExpiration(context),
	}
}

func (d Dispatcher) CheckExpirationAndDispatch() {
	if d.expiration.Expired() {
		d.mailer.Dispatch()
	}
}
