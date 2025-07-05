package common

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

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

func (t TokenType) String() string {
	return string(t)
}
