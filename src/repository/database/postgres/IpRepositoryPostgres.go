package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"

	co "github.com/Nikitarsis/goTokens/common"
	in "github.com/Nikitarsis/goTokens/repository/interfaces"
	ad "github.com/Nikitarsis/goTokens/repository/database/postgres/adapter"
)

// IpRepositoryPostgres реализует интерфейс IIpRepository
type IpRepositoryPostgres struct {
	adapter ad.IAdapterSQL
}

// NewIpRepositoryPostgres создает новый экземпляр IpRepositoryPostgres
func NewIpRepositoryPostgres(config in.IPostgresConfig) in.IIpRepository {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		return nil
	}
	adapter := ad.CreateAdapterSQL(db)
	return &IpRepositoryPostgres{adapter: adapter}
}

// SaveIp сохраняет IP-адрес в репозитории
func (ir IpRepositoryPostgres) SaveIp(ip co.DataIP) error {
	ir.adapter.AddIp(ip)
	return nil
}

// CheckIp проверяет, существует ли IP-адрес в репозитории
func (ir IpRepositoryPostgres) CheckIp(ip co.DataIP) bool {
	return ir.adapter.CheckIp(ip)
}