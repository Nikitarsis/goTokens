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
		// Проверка наличия ключа в токене
		if !ok {
			return nil, co.ErrInvalidToken
		}
		//Попытка парсинга ключа из токена
		kid, err := co.GetUUIDFromString(kidRaw)
		if err != nil {
			return nil, co.ErrInvalidToken
		}
		// Поиск ключа в репозитории
		key, ok := tp.keyGetter(kid)
		if !ok {
			return nil, co.ErrInvalidToken
		}
		return key.GetValue(), nil
	})
}

// GetTokenData извлекает данные токена из строки токена
// Если токен невалиден, возвращает ошибку без значения co.TokenData
func (tp tokenParser) GetTokenData(tokenRaw co.Token) (co.TokenData, error) {
	// Парсинг токена
	token, err := tp.parseToken(tokenRaw)
	if err != nil {
		return co.TokenData{}, err
	}
	//Проверка токена, невалидный токен НЕ возвращается как co.TokenData
	if !token.Valid {
		return co.TokenData{}, co.ErrInvalidToken
	}
	claims := token.Claims.(jwt.MapClaims)
	// Извлечение данных из токена
	sub, ok := claims["sub"].(string)
	if !ok {
		return co.TokenData{}, co.ErrInvalidToken
	}
	userId, err := co.GetUUIDFromString(sub)
	if err != nil {
		return co.TokenData{}, err
	}
	kid, ok := claims["kid"].(string)
	if !ok {
		return co.TokenData{}, co.ErrInvalidToken
	}
	keyId, err := co.GetUUIDFromString(kid)
	if err != nil {
		return co.TokenData{}, err
	}
	tokType, ok := claims["type"].(string)
	if !ok {
		return co.TokenData{}, co.ErrInvalidToken
	}
	tokenType := co.GetTokenType(tokType)
	jti, ok := claims["jti"].(string)
	if !ok {
		return co.TokenData{}, co.ErrInvalidToken
	}
	tokenId, err := co.GetUUIDFromString(jti)
	if err != nil {
		return co.TokenData{}, err
	}
	return co.TokenData{
		Token:   tokenRaw,
		TokenId: tokenId,
		UserId:  userId,
		KeyId:   keyId,
		Type:    tokenType,
	}, nil
}
