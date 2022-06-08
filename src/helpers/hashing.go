package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasword(pass string) (string, error) {
	hassPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hassPass), nil
}

func CheckPassword(hassPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hassPass), []byte(password))

	if err != nil {
		return false
	}

	return true
}
