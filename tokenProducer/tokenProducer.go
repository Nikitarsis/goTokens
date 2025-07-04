package tokenProducer

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UUID [16]byte

/*Создание токена*/
type tokenProducer struct {
	secretKey []byte // секретный ключ
    issuer string // издатель токена
    jtiSupplier func() UUID // функция для генерации уникального идентификатора токена
}

func NewTokenProducer(secretKey, issuer string, jtiSupplier func() UUID) *tokenProducer {
	return &tokenProducer{
		secretKey:  []byte(secretKey),
		issuer:     issuer,
		jtiSupplier: jtiSupplier,
	}
}

/*Создание тела токена*/
func (tp *tokenProducer) createClaims(tokenType string, uid UUID, groupId UUID) jwt.MapClaims {
    return jwt.MapClaims{
        "type": tokenType,//тип токена: authorized или refresh
        "group": groupId,// ID пары токенов
        "jti": tp.jtiSupplier(),// уникальный идентификатор токена
        "iat": time.Now().Unix(),// время создания токена
        "sub": uid,// ID пользователя
        "iss": tp.issuer,// издатель токена
    }
}

/*Создание токена определённого типа*/
func (tp *tokenProducer) CreateToken(uid UUID, groupId UUID, tokenType string) (UUID, string, error) {
    jti := tp.jtiSupplier()
    token := jwt.NewWithClaims(jwt.SigningMethodES512, tp.createClaims(tokenType, uid, groupId))
    signedString, err := token.SignedString(tp.secretKey)
    return jti, signedString, err
}

func (tp *tokenProducer) CreateAuthorizedToken(uid UUID, groupId UUID) (UUID, string, error) {
    return tp.CreateToken(uid, groupId, "authorized")
}

func (tp *tokenProducer) CreateRefreshToken(uid UUID, groupId UUID) (UUID, string, error) {
    return tp.CreateToken(uid, groupId, "refresh")
}
