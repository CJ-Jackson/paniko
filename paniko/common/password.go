//go:generate mockgen -write_package_comment=false -package=common -source=password.go -destination=password.mock.go
//go:generate debugflag password.mock.go

package common

import (
	"crypto/sha512"
	"encoding/base64"
)

type Password interface {
	SaltPassword(password string) string
	CheckPassword(password, hash string) bool
}

type password struct {
	salt string
}

func NewPassword(salt string) Password {
	return password{
		salt: salt,
	}
}

func (p password) SaltPassword(password string) string {
	hash := sha512.New()
	hash.Write([]byte(p.salt))
	hash.Write([]byte(password))

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}

func (p password) CheckPassword(password, hash string) bool {
	return hash == p.SaltPassword(password)
}
