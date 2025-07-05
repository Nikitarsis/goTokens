package common

import (
	"encoding/base64"
)

type UUID struct {
	value [16]byte
}

func (u UUID) ToString() string {
	return base64.RawStdEncoding.EncodeToString(u.value[:])
}

func GetUUIDFromString(s string) (UUID, error) {
	var u UUID
	data, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, err
	}
	copy(u.value[:], data)
	return u, nil
}