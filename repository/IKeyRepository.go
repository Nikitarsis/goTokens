package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// Интерфейс для репозитория ключей
type IKeyRepository interface {
	co.IKeyKeepingRepository
	IKeyController
}
