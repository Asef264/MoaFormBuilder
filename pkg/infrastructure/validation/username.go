package validation

import (
	errHnadler "moaformbuilder/pkg/infrastructure/errHandler"
	"regexp"
)

// a valid username must be in English
func ValidateUseranme(username string) error {
	r, _ := regexp.Compile(`[a-zA-Z]`)
	if !r.MatchString(username) {
		newErr := errHnadler.NewMyError("username contains not english letters", 403) // all validation errors corresponded with a 403 HTTP status code.
		newErr.WithDetail("username validation")
		return newErr.MainError
	}
	return nil
}
