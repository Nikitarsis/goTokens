package tokenProducer

import (
	co "github.com/Nikitarsis/goTokens/common"
	"github.com/dgrijalva/jwt-go"
)

// tokenParser - парсер токенов
type tokenParser struct {
	keyGetter func(co.UUID) (co.Key, bool)
}

// NewTokenParser создает новый экземпляр tokenParser
func NewTokenParser(secretKeyProducer func(co.UUID) (co.Key, bool)) *tokenParser {
	return &tokenParser{
		keyGetter: secretKeyProducer,
	}
}

// parseToken парсит токен и возвращает его
func (tp tokenParser) parseToken(token co.Token) (*jwt.Token, error) {
	return jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
		kidRaw, ok := token.Claims.(jwt.MapClaims)["kid"].(string)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}
		kid, err := co.GetUUIDFromString(kidRaw)
		if err != nil {
			return nil, jwt.ErrInvalidKey
		}
		key, ok := tp.keyGetter(kid)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}
		return key.GetValue(), nil
	})
}

// GetTokenData извлекает данные токена из строки токена
func (tp tokenParser) GetTokenData(tokenStr co.Token) (co.TokenData, error) {
	token, err := tp.parseToken(tokenStr)
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
