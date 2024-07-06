package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPass(hashedPass, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)); err != nil {
		return err
	}
	return nil
}
