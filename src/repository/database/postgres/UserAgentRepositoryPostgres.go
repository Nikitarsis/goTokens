package postgres

import (
	"database/sql"

	co "github.com/Nikitarsis/goTokens/common"
	ad "github.com/Nikitarsis/goTokens/repository/database/postgres/adapter"
	inter "github.com/Nikitarsis/goTokens/repository/interfaces"
)

// Postgres репозиторий агентов пользователей
type UserAgentRepositoryPostgres struct {
	adapter ad.IAdapterSQL
}

// CreatePostgresUserAgentRepository создает новый экземпляр UserAgentRepositoryPostgres.
func CreatePostgresUserAgentRepository(config inter.IPostgresConfig) co.IUserAgentRepository {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		return nil
	}
	adapter := ad.CreateAdapterSQL(db)
	return &UserAgentRepositoryPostgres{
		adapter: adapter,
	}
}

// SaveUserAgent сохраняет агента пользователя в репозитории.
func (r *UserAgentRepositoryPostgres) SaveUserAgent(kid co.UUID, agent co.UserAgentData) error {
	r.adapter.AddUserAgent(kid, agent.Data)
	return nil
}

// GetUserAgent загружает агента пользователя из репозитория.
func (r *UserAgentRepositoryPostgres) CheckUserAgent(id co.UUID, userAgent co.UserAgentData) bool {
	ret, err := r.adapter.GetUserAgent(id)
	if err != nil {
		return false
	}
	return ret == userAgent.Data
}