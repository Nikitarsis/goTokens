package tokenProducer

import (
	"github.com/dgrijalva/jwt-go"
)

type tokenChecker struct {
	secretKey []byte
	issuer    string
}

func NewTokenChecker(secretKey, issuer string) *tokenChecker {
	return &tokenChecker{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}

func (tc *tokenChecker) CheckToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tc.secretKey, nil
	})
}
