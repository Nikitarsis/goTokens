package common

// TokenData - структура, представляющая данные токена
type TokenData struct {
	Token   Token
	TokenId UUID
	UserId  UUID
	KeyId   UUID
	Type    TokenType
}
