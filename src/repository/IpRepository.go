package repository

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
	pg "github.com/Nikitarsis/goTokens/repository/database/postgres"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// IpRepository реализует интерфейс IIpRepository
//
// Он содержит методы горячего хэширования, позволяющие быстро сохранять и проверять IP-адреса.
type IpRepository struct {
	db             in.IIpRepository
	hotCacheToSave *co.SafeMap[string, co.DataIP]
	tracePort      bool
}

// NewIpRepository создает новый экземпляр IpRepository
func NewIpRepository(config in.IRepositoryConfig) in.IIpRepository {
	db := pg.NewIpRepositoryPostgres(config)
	hotCacheToSave := co.CreateSafeMap[string, co.DataIP]()
	return &IpRepository{
		db:             db,
		hotCacheToSave: hotCacheToSave,
		tracePort:      config.TracePorts(),
	}
}

// SaveIp сохраняет IP-адрес в репозитории
func (ir IpRepository) SaveIp(ip co.DataIP) error {
	var ipKey string
	// проверка условия хранения
	if ir.tracePort {
		ipKey = fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port)
	} else {
		ipKey = ip.IP.String()
	}
	// сохранение в кэш и назначение удаления
	ir.hotCacheToSave.Store(ipKey, ip)
	defer ir.hotCacheToSave.Delete(ipKey)
	// сохранение в БД
	return ir.db.SaveIp(ip)
}

// CheckIp проверяет, существует ли IP-адрес в репозитории
func (ir IpRepository) CheckIp(ip co.DataIP) bool {
	var ipKey string
	// проверка условия хранения
	if ir.tracePort {
		ipKey = fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port)
	} else {
		ipKey = ip.IP.String()
	}
	// проверка наличия в кэше, если есть, возвращается true
	_, check := ir.hotCacheToSave.Load(ipKey)
	if check {
		return true
	}
	// проверка наличия в БД
	return ir.db.CheckIp(ip)
}
