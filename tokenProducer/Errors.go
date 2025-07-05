package tokenProducer

import "errors"

var (
	// ErrInvalidToken - Ошибка, возникающая при недействительном токене (!Valid)
	ErrInvalidToken = errors.New("invalid token")
)
