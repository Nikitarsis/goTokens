package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// Интерфейс для управления ключами
type IKeyController interface {
	DropKey(kid co.UUID) bool
}
