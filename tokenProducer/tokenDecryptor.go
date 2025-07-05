package tokenProducer

import (
	co "github.com/Nikitarsis/goTokens/common"
	"github.com/dgrijalva/jwt-go"
)

type tokenChecker struct {
	keyGetter func(co.UUID) (co.Key, bool)
}

func NewTokenChecker(secretKeyProducer func(co.UUID) (co.Key, bool)) *tokenChecker {
	return &tokenChecker{
		keyGetter: secretKeyProducer,
	}
}

func (tc *tokenChecker) parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		kidRaw, ok := token.Claims.(jwt.MapClaims)["kid"].(string)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}
		kid, err := co.GetUUIDFromString(kidRaw)
		if err != nil {
			return nil, jwt.ErrInvalidKey
		}
		key, ok := tc.keyGetter(kid)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}
		return key, nil
	})
}

func (tc *tokenChecker) GetTokenData(tokenString string) (co.TokenData, error) {
	token, err := tc.parseToken(tokenString)
	if err != nil {
		return co.TokenData{}, err
	}
	if !token.Valid {
		return co.TokenData{}, ErrInvalidToken
	}
	userId, err := co.GetUUIDFromString(token.Claims.(jwt.MapClaims)["sub"].(string))
	if err != nil {
		return co.TokenData{}, err
	}
	keyId, err := co.GetUUIDFromString(token.Claims.(jwt.MapClaims)["kid"].(string))
	if err != nil {
		return co.TokenData{}, err
	}
	tokenType := co.GetTokenType(token.Claims.(jwt.MapClaims)["type"].(string))
	return co.TokenData{
		UserId: userId,
		KeyId:  keyId,
		Type:   tokenType,
	}, nil
}
