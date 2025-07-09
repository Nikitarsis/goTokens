package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"

	co "github.com/Nikitarsis/goTokens/common"
	ad "github.com/Nikitarsis/goTokens/repository/database/postgres/adapter"
	inter "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// Postgres репозиторий ключей
type KeyRepositoryPostgres struct {
	adapter ad.IAdapterSQL
}

// CreatePostgresKeyRepository создает новый экземпляр KeyRepositoryPostgres
func CreatePostgresKeyRepository(config inter.IPostgresConfig) inter.IKeyRepository {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		return nil
	}
	adapter := ad.CreateAdapterSQL(db)
	return &KeyRepositoryPostgres{
		adapter: adapter,
	}
}

// SaveKey сохраняет ключ в репозитории.
func (r *KeyRepositoryPostgres) SaveKey(key co.Key) {
	r.adapter.AddKey(key.GetKid(), key)
}

// GetKey загружает ключ из репозитория.
func (r *KeyRepositoryPostgres) GetKey(kid co.UUID) (co.Key, bool) {
	return r.adapter.GetKey(kid)
}

// DropKey удаляет ключ из репозитория.
func (r *KeyRepositoryPostgres) DropKey(kid co.UUID) bool {
	return true
}
