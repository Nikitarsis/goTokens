package common

import (
	"encoding/base64"
	"crypto/rand"
)

// Key - структура, представляющая ключ
type Key struct {
	kid   UUID
	value []byte
}

// CreateTestKey создает тестовый ключ
func CreateTestKey() Key {
	key := make([]byte, 64)
	rand.Read(key) // Случайная генерация ключа
	return Key{kid: GetTestUUID(), value: key}
}

// CreateKeyFromBytes создает новый ключ из массива байтов
func CreateKeyFromBytes(kid UUID, b []byte) Key {
	return Key{kid: kid, value: b}
}

// CreateKeyFromString создает новый ключ из строки типа base64
func CreateKeyFromString(kid UUID, s string) (Key, error) {
	ret, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Key{}, err
	}
	return Key{kid: kid, value: ret}, nil
}

// GetKid возвращает идентификатор ключа
func (k Key) GetKid() UUID {
	return k.kid
}

// GetValue возвращает значение ключа
func (k Key) GetValue() []byte {
	return k.value
}

// ToString возвращает строковое представление ключа в формате base64
func (k Key) ToString() string {
	return base64.StdEncoding.EncodeToString(k.value)
}
