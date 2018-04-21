package common

import (
	"log"
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
