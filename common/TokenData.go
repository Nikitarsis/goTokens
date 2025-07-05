package common

// TokenData - структура, представляющая данные токена
type TokenData struct {
	UserId UUID
	KeyId  UUID
	Type   TokenType
}
