package utils

import "golang.org/x/crypto/bcrypt"

type Password string

func (p Password) Encode() (string, error) {
	pwd, err := bcrypt.GenerateFromPassword(p.Byte(), bcrypt.DefaultCost)
	return string(pwd), err
}

func (p Password) Compare(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), p.Byte())
	return err
}

func (p Password) Byte() []byte {
	return []byte(p)
}
