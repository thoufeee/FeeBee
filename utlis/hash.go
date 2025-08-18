package utlis

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// password check
func PasswordStrength(pass string) bool {
	if len(pass) < 6 {
		return false
	}

	var hasupper, hasdigit, hasSpecial bool

	for _, ch := range pass {
		switch {
		case unicode.IsUpper(ch):
			hasupper = true
		case unicode.IsDigit(ch):
			hasdigit = true
		case unicode.IsPunct(ch), unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasupper && hasdigit && hasSpecial
}

// hashing
func GenerateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(hash), err
}

// check pass
func CheckPass(pass, newpass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(newpass))
	return err == nil
}
