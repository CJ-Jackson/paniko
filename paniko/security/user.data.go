package security

import (
	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/uri"
)

type User struct {
	username string
}

func (u User) Username() string { return u.username }

func (u User) CheckIfUser() {}

func (u User) CheckIfGuest() { common.HaltForbidden("User is already logged in.") }

type Guest struct{}

func (g Guest) Username() string { return "Guest" }

func (g Guest) CheckIfUser() { common.HaltSeeOther(uri.Login) }

func (g Guest) CheckIfGuest() {}
