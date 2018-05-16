package common

import (
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

func GetMuxer(context ctx.BackgroundContext) *httprouter.Router {
	return context.Persist("muxer-148a5db092ec73a35160dcb667a1fb45", func() (interface{}, error) {
		muxer := httprouter.New()

		return muxer, nil
	}).(*httprouter.Router)
}
