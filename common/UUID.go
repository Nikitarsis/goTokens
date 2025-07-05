package common

import (
	"encoding/base64"
	"errors"
)

type UUID struct {
	value [16]byte
}

func (u UUID) ToString() string {
	return base64.RawStdEncoding.EncodeToString(u.value[:])
}

func GetUUIDFromBytes(b []byte) (UUID, error) {
	var u UUID
	if len(b) != 16 {
		return UUID{}, errors.New("invalid byte length")
	}
	copy(u.value[:], b)
	return u, nil
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