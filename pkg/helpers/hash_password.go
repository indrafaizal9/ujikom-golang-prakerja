package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	salt := 4
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	fmt.Println(hash)
	return string(hash)
}

func ComparePassword(hashedPassword, password string) bool {
	hash, pass := []byte(hashedPassword), []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
