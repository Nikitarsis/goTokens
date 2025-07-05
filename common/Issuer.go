package common

type Issuer struct {
	value string
}

func (i Issuer) String() string {
	return i.value
}

func NewIssuer(value string) Issuer {
	return Issuer{value: value}
}
