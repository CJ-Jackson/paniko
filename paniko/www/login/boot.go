package login

import (
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/cjtoolkit/ctx"
)

func Boot(context ctx.BackgroundContext) {
	muxer := common.GetMuxer(context)
	bootLogin(NewLoginController(context), muxer)
}
