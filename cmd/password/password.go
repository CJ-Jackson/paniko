package main

import (
	"os"

	"github.com/CJ-Jackson/paniko/paniko/security"
	"github.com/cjtoolkit/ctx"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	context := ctx.NewBackgroundContext()
	controller := security.GetUserController(context)

	controller.UpdateUser("admin", os.Args[1])
	controller.Save()
}
