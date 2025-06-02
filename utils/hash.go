package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),19)
	return string(bytes),err
}
func CompareHashPassword(hashedPassword,plainPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainPassword))
}