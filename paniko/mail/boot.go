package mail

import "github.com/cjtoolkit/ctx"

func Boot(context ctx.BackgroundContext) {
	bootDispatcher(NewDispatcher(context))
}
