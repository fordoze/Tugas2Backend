package tools

import (
	"golang.org/x/crypto/bcrypt"
)

// function genarate password Parameter stirng
func GenaratePassword(password string) string {

	passwordByte := []byte(password)

	hash, _ := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)

	stringPassword := string(hash)

	return stringPassword
}
