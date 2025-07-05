package common

// TokensPair - DAO структура, представляющая пару токенов
type TokensPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
