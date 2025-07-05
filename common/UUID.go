package common

import (
	"encoding/base64"
)

type UUID [16]byte

func (u UUID) ToString() string {
	return base64.RawStdEncoding.EncodeToString(u[:]);
}

func GetUUIDFromString(s string) (UUID, error) {
	var u UUID
	data, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return u, err
	}
	copy(u[:], data)
	return u, nil
}