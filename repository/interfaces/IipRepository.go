package repository

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type IIpRepository interface {
	SaveIp(kid co.UUID, ip co.DataIP) error
	CheckIp(kid co.UUID, ip co.DataIP) bool
}