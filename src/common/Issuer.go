package common

// Issuer - структура, представляющая издателя токена
type Issuer struct {
	value string
}

// String возвращает строковое представление издателя токена
func (i Issuer) String() string {
	return i.value
}

// NewIssuer создает нового издателя токена
func NewIssuer(value string) Issuer {
	return Issuer{value: value}
}
