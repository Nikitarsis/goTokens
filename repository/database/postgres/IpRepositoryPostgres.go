package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"

	co "github.com/Nikitarsis/goTokens/common"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
	ad "github.com/Nikitarsis/goTokens/repository/database/postgres/adapter"
)

type IpRepositoryPostgres struct {
	adapter ad.IAdapterSQL
}

func NewIpRepositoryPostgres(config in.IPostgresConfig) in.IIpRepository {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		return nil
	}
	adapter, err := ad.CreateAdapterSQL(db)
	if err != nil {
		return nil
	}
	return &IpRepositoryPostgres{adapter: adapter}
}

func (ir IpRepositoryPostgres) SaveIp(ip co.DataIP) error {
	ir.adapter.AddIp(ip)
	return nil
}

func (ir IpRepositoryPostgres) CheckIp(ip co.DataIP) bool {
	return ir.adapter.CheckIp(ip)
}