//go:generate mockgen -write_package_comment=false -package=common -source=error.go -destination=error.mock.go
//go:generate debugflag error.mock.go

package common

import (
	"log"
	"net/http"
	"os"

	"github.com/CJ-Jackson/ctx"
)

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
	const name = "error-service-54dedf6100fb4eecee111fe005cff343"
	if _ErrorService, ok := context.Ctx(name).(ErrorService); ok {
		return _ErrorService
	}

	_ErrorService := errorService{
		log: log.New(os.Stderr, "INFO: ", log.Lshortfile),
	}

	context.SetCtx(name, _ErrorService)
	return _ErrorService
}

type NoError struct{}

func Halt() {
	panic(NoError{})
}

type HttpError struct {
	Code    int
	Message string
}

func HaltNotFound(message string) {
	panic(HttpError{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func HaltForbidden(message string) {
	panic(HttpError{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

func HaltInternalServerError(message string) {
	panic(HttpError{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func HaltCustomError(code int, message string) {
	panic(HttpError{
		Code:    code,
		Message: message,
	})
}

type HttpRedirectError struct {
	Code     int
	Location string
}

func HaltMovedPermanently(location string) {
	panic(HttpRedirectError{
		Code:     http.StatusMovedPermanently,
		Location: location,
	})
}

func HaltSeeOther(location string) {
	panic(HttpRedirectError{
		Code:     http.StatusSeeOther,
		Location: location,
	})
}
