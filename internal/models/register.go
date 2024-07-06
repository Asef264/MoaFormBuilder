package models

import (
	"errors"
	"moaformbuilder/pkg/infrastructure/validation"

	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r Register) Validate() error {
	if len(r.Name) == 0 || len(r.Name) < 3 || len(r.Name) > 30 {
		return errors.New("name lengh not valid")
	}
	if err := validation.ValidateUseranme(r.Username); err != nil {
		return err
	}
	return nil
}
