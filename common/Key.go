package common

import (
	"encoding/base64"
)

type Key struct {
	kid   UUID
	value []byte
}

func CreateKeyFromBytes(kid UUID, b []byte) Key {
	return Key{kid: kid, value: b}
}

func CreateKeyFromString(kid UUID, s string) (Key, error) {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Key{}, err
	}
	return Key{kid: kid, value: ret}, nil
}

func (k Key) GetKid() UUID {
	return k.kid
}

func (k Key) ToString() string {
	return base64.StdEncoding.EncodeToString(k.value)
}