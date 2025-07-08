package common

// TokenData - структура, представляющая данные токена
type TokenData struct {
	Token   Token
	TokenId UUID
	UserId  UUID
	KeyId   UUID
	Type    TokenType
}

// GetTestTokenData - возвращает тестовые данные токена
func GetTestTokenData(tokenType TokenType) TokenData {
	return TokenData{
		Token:   GetTestToken(),
		TokenId: GetTestUUID(),
		UserId:  GetTestUUID(),
		KeyId:   GetTestUUID(),
		Type:    tokenType,
	}
}