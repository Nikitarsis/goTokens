package common

// Token - структура-обёртка для хранения токена
type Token struct {
	// Value - значение токена
	Value string
}

// ToString - возвращает строковое представление токена
func (t Token) ToString() string {
	return t.Value
}

// GetTestToken - возвращает тестовый токен
func GetTestToken() Token {
	return Token{
		Value: "just.test.token",
	}
}