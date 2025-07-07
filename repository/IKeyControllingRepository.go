package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type IKeyController interface {
	DropKey(kid co.UUID) bool
}
