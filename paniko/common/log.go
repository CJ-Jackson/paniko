//go:generate mockgen -write_package_comment=false -package=common -source=log.go -destination=log.mock.go
//go:generate debugflag log.mock.go

package common

type Logger interface {
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
