package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type IKeyRepository interface {
	co.IKeyKeepingRepository
	IKeyController
}
