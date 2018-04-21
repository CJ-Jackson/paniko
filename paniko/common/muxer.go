package common

import (
	"github.com/CJ-Jackson/ctx"
	"github.com/julienschmidt/httprouter"
)

func GetMuxer(context ctx.BackgroundContext) *httprouter.Router {
	name := "muxer-148a5db092ec73a35160dcb667a1fb45"
	if muxer, ok := context.Ctx(name).(*httprouter.Router); ok {
		return muxer
	}

	muxer := httprouter.New()

	context.SetCtx(name, muxer)
	return muxer
}
