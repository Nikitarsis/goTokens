package repository

import (
	"fmt"

	co "github.com/Nikitarsis/goTokens/common"
	pg "github.com/Nikitarsis/goTokens/repository/database/postgres"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
)

type IpRepository struct {
	db             in.IIpRepository
	hotCacheToSave *co.SafeMap[string, co.DataIP]
	tracePort      bool
}

func NewIpRepository(config in.IRepositoryConfig) in.IIpRepository {
	db := pg.NewIpRepositoryPostgres(config)
	hotCacheToSave := &co.SafeMap[string, co.DataIP]{}
	return &IpRepository{
		db:              db,
		hotCacheToSave:  hotCacheToSave,
		tracePort:      config.TracePorts(),
	}
}

func (ir IpRepository) SaveIp(ip co.DataIP) error {
	var ipKey string
	if ir.tracePort {
		ipKey = fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port)
	} else {
		ipKey = ip.IP.String()
	}
	ir.hotCacheToSave.Store(ipKey, ip)
	defer ir.hotCacheToSave.Delete(ipKey )
	return ir.db.SaveIp(ip)
}

func (ir IpRepository) CheckIp(ip co.DataIP) bool {
	var ipKey string
	if ir.tracePort {
		ipKey = fmt.Sprintf("%s:%d", ip.IP.String(), ip.Port)
	} else {
		ipKey = ip.IP.String()
	}
	_, check := ir.hotCacheToSave.Load(ipKey)
	if check {
		return true
	}
	return ir.db.CheckIp(ip)
}