package tokenProducer

import (
	"time"

	co "github.com/Nikitarsis/goTokens/common"
	"github.com/dgrijalva/jwt-go"
)

// Создание токена
type tokenProducer struct {
	issuer      co.Issuer      // издатель токена
	jtiSupplier func() co.UUID // функция для генерации уникального идентификатора токена
}

// NewTokenProducer создает новый экземпляр tokenProducer
func NewTokenProducer(issuer co.Issuer, jtiSupplier func() co.UUID) *tokenProducer {
	return &tokenProducer{
		issuer:      issuer,
		jtiSupplier: jtiSupplier,
	}
}

// Создание тела токена
func (tp tokenProducer) createClaims(tokenType co.TokenType, uid co.UUID, keyId co.UUID) jwt.MapClaims {
	return jwt.MapClaims{
		"type": tokenType.String(), //тип токена: authorized или refresh
		"kid":  keyId.ToString(),   // ID ключа шифрования
		"jti":  tp.jtiSupplier(),   // уникальный идентификатор токена
		"iat":  time.Now().Unix(),  // время создания токена
		"sub":  uid.ToString(),     // ID пользователя
		"iss":  tp.issuer.String(), // издатель токена
	}
}

// Создание токена определённого типа
func (tp tokenProducer) createToken(key co.Key, uid co.UUID, tokenType co.TokenType) (co.UUID, string, error) {
	jti := tp.jtiSupplier()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, tp.createClaims(tokenType, uid, key.GetKid()))
	signedString, err := token.SignedString(key.GetValue())
	return jti, signedString, err
}

// CreateAccessToken создает новый access токен
func (tp tokenProducer) CreateAccessToken(key co.Key, uid co.UUID) (co.UUID, string, error) {
	return tp.createToken(key, uid, co.AccessToken)
}

// CreateRefreshToken создает новый refresh токен
func (tp tokenProducer) CreateRefreshToken(key co.Key, uid co.UUID) (co.UUID, string, error) {
	return tp.createToken(key, uid, co.RefreshToken)
}
