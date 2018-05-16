package common

import (
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

func GetMuxer(context ctx.BackgroundContext) *httprouter.Router {
	const name = "muxer-148a5db092ec73a35160dcb667a1fb45"
	if muxer, ok := context.Get(name).(*httprouter.Router); ok {
		return muxer
	}

	muxer := httprouter.New()

	context.Set(name, muxer)
	return muxer
}
