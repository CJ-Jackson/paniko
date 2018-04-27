package login

import (
	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

func Boot(context ctx.BackgroundContext) {
	muxer := common.GetMuxer(context)
	bootLogin(NewLoginController(context), muxer)
}
