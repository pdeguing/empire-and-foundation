package data

import (
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks if the given password matches the user's current password.
func CheckPassword(userPassword string, password string) (ok bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
