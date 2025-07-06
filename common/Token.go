package common

type Token struct {
	Value string	`json:"token"`
}

func (t Token) ToString() string {
	return t.Value
}