package tokenProducer

import "errors"

var (
	// Ошибка, возникающая при недействительном токене (!Valid)
	ErrInvalidToken = errors.New("invalid token")
)
