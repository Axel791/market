package domains

import (
	"fmt"
	"regexp"
)

const (
	ValidatePasswordError = "password validation failed"
	ValidateLoginError    = "login validation failed"
)

type User struct {
	ID       int64
	Login    string
	Password string
}

func (u *User) ValidateLogin() error {
	if u.Login == "" {
		return fmt.Errorf(ValidateLoginError)
	}
	latinRegex := regexp.MustCompile(`^[A-Za-z]+$`)
	if !latinRegex.MatchString(u.Login) {
		return fmt.Errorf(ValidateLoginError)
	}

	return nil
}

func (u *User) ValidatePassword() error {
	if len(u.Password) < 6 {
		return fmt.Errorf(ValidatePasswordError)
	}
	return nil
}
