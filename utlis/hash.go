package utlis

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var passwordRegex *regexp.Regexp

// regex password set
func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("failed to load env")
	}

	regexPattern := os.Getenv("PASSWORD_REGEX")
	if regexPattern == "" {
		log.Println("regex invalid")
	}

	passwordRegex, err = regexp.Compile(regexPattern)

	if err != nil {
		log.Fatalln("invalid password regex")
	}

}

// password check
func PasswordStrength(pass string) bool {
	if len(pass) < 6 {
		return false
	}

	if !passwordRegex.MatchString(pass) {
		return false
	}

	return true
}
