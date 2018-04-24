package common

import (
	"log"

	"os"

	"github.com/CJ-Jackson/ctx"
)

func CheckErrorAndPanic(err error) {
	if nil != err {
		log.Panic(err)
	}
}

func CheckErrorAndExit(err error) {
	if nil != err {
		log.Fatal(err)
	}
}

type ErrorService interface {
	CheckErrorAndPanic(err error)
	CheckErrorAndLog(err error)
}

type errorService struct {
	log Logger
}

func (e errorService) CheckErrorAndPanic(err error) {
	if nil != err {
		e.log.Panic(err)
	}
}

func (e errorService) CheckErrorAndLog(err error) {
	if nil != err {
		e.log.Print(err)
	}
}

func GetErrorService(context ctx.BackgroundContext) ErrorService {
	name := "error-service-54dedf6100fb4eecee111fe005cff343"
	if _ErrorService, ok := context.Ctx(name).(ErrorService); ok {
		return _ErrorService
	}

	_ErrorService := errorService{
		log: log.New(os.Stderr, "INFO: ", log.Lshortfile),
	}

	context.SetCtx(name, _ErrorService)
	return _ErrorService
}
