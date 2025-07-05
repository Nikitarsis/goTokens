package common

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func (t TokenType) String() string {
	return string(t)
}
