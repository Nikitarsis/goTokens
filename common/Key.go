package common

import (
	"encoding/base64"
)

type Key struct {
	value []byte
}

func ParseFromString(s string) (Key, error) {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Key{}, err
	}
	return Key{value: ret}, nil
}

func (k Key) ToString() string {
	return base64.StdEncoding.EncodeToString(k.value)
}