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
	salt         string
	errorService ErrorService
}

func NewPassword(salt string, errorService ErrorService) Password {
	return password{
		salt:         salt,
		errorService: errorService,
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
