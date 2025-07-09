package interfaces

import (
	co "github.com/Nikitarsis/goTokens/common"
)

type IIpRepository interface {
	SaveIp(ip co.DataIP) error
	CheckIp(ip co.DataIP) bool
}