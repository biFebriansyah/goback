package helpers

import (
	"fmt"

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
	fmt.Println(hassPass)
	fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(hassPass), []byte(password))
	fmt.Println(err)

	if err != nil {
		return false
	}

	return true
}
