package postgres

import (
	co "github.com/Nikitarsis/goTokens/common"
)

// IAdapterSQL определяет интерфейс для работы с адаптером SQL
type IAdapterSQL interface {
	CreateTablesIFNotExists()
	AddKey(kid co.UUID, key co.Key)
	GetKey(kid co.UUID) (co.Key, bool)
	RemoveKey(kid co.UUID)
	AddUserAgent(kid co.UUID, userAgent string)
	GetUserAgent(kid co.UUID) (string, error)
	AddIp(ip co.DataIP)
	CheckIp(ip co.DataIP) bool
}