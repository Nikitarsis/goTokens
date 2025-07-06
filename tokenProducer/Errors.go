package tokenProducer

import "errors"

var (
	// ErrInvalidToken - Ошибка, возникающая при недействительном токене (!Valid)
	ErrInvalidToken = errors.New("invalid token")
	// ErrNoFindKey - Ошибка, возникающая при отсутствии ключа
	ErrNoFindKey   = errors.New("key not found")
)
