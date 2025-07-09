package adapter

import (
	"database/sql"
	"embed"
)

//go:embed scripts
var sqlScripts embed.FS

// CreateAdapterSQL создает новый экземпляр адаптера SQL
//
// Парсит включённую директорию и загружает SQL-скрипты, если что-то не срабатывает, поднимает панику.
func CreateAdapterSQL(db *sql.DB) IAdapterSQL {
	CreateTablesQuery, err := sqlScripts.ReadFile("scripts/CreateTables.sql")
	if err != nil {
		panic(err)
	}
	AddKeyQuery, err := sqlScripts.ReadFile("scripts/AddKey.sql")
	if err != nil {
		panic(err)
	}
	GetKeyQuery, err := sqlScripts.ReadFile("scripts/GetKey.sql")
	if err != nil {
		panic(err)
	}
	RemoveKeyQuery, err := sqlScripts.ReadFile("scripts/RemoveKey.sql")
	if err != nil {
		panic(err)
	}
	AddUserAgentsQuery, err := sqlScripts.ReadFile("scripts/AddUserAgent.sql")
	if err != nil {
		panic(err)
	}
	GetUserAgentsQuery, err := sqlScripts.ReadFile("scripts/GetUserAgent.sql")
	if err != nil {
		panic(err)
	}
	AddIpQuery, err := sqlScripts.ReadFile("scripts/AddIP.sql")
	if err != nil {
		panic(err)
	}
	CheckIpQuery, err := sqlScripts.ReadFile("scripts/CheckIP.sql")
	if err != nil {
		panic(err)
	}

	return &adapterSQL{
		Exec: func(query string, args ...interface{}) (sql.Result, error) {
			return db.Exec(query, args...)
		},
		Query: func(query string, args ...interface{}) (*sql.Rows, error) {
			return db.Query(query, args...)
		},
		CreateTablesQuery:  string(CreateTablesQuery),
		AddKeyQuery:        string(AddKeyQuery),
		RemoveKeyQuery:     string(RemoveKeyQuery),
		AddUserAgentsQuery: string(AddUserAgentsQuery),
		GetUserAgentsQuery: string(GetUserAgentsQuery),
		AddIpQuery:         string(AddIpQuery),
		CheckIpQuery:       string(CheckIpQuery),
		GetKeyQuery:        string(GetKeyQuery),
	}
}
