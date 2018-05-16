package www

import (
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/cjtoolkit/ctx"
)

func Boot(context ctx.BackgroundContext) {
	muxer := common.GetMuxer(context)
	homeBoot(NewHomeController(context), muxer)
}
