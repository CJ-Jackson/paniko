package mail

import "github.com/CJ-Jackson/ctx"

func Boot(context ctx.BackgroundContext) {
	bootDispatcher(NewDispatcher(context))
}
