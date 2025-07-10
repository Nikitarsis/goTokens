package interfaces

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// IIpRepository определяет интерфейс репозитория для работы с IP-адресами
type IIpRepository interface {
	SaveIp(ip co.DataIP) error
	CheckIp(ip co.DataIP) bool
}