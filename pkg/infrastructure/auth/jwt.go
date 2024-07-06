package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	WantedClaim interface{}
}

func (c *Claims) GenerteToken() (string, error) {
	secretKey := []byte("secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	return token.SignedString(secretKey)
}

func ParseToken(token string) (Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(tkn *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if !t.Valid {
		return Claims{}, errors.New("invalid token ")
	}
	claim, ok := t.Claims.(*Claims)
	if !ok {
		return Claims{}, errors.New("invalid claims")
	}
	return *claim, nil

}
