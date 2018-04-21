package paniko

import (
	"fmt"
	"net/http"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/common"
)

func Boot() {
	context := ctx.NewBackgroundContext()

	{
		muxer := common.GetMuxer(context)

		fmt.Println("Running Server at :8080")
		http.ListenAndServe(":8080", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			req, _ = ctx.NewContext(req, res)
			muxer.ServeHTTP(res, req)
		}))
	}
}
