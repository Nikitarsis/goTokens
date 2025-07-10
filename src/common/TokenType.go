package common

// TokenType - тип токена
type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

// GetTokenType возвращает тип токена по его строковому представлению
func GetTokenType(s string) TokenType {
	switch s {
	case "access":
		return AccessToken
	case "refresh":
		return RefreshToken
	default:
		return ""
	}
}

// String возвращает строковое представление типа токена
func (t TokenType) String() string {
	return string(t)
}
