package postgres

import (
	"database/sql"
	"embed"
)

//go:embed scripts
var sqlScripts embed.FS

func CreateAdapterSQL(db *sql.DB) (IAdapterSQL, error) {
	CreateTablesQuery, err := sqlScripts.ReadFile("scripts/CreateTables.sql")
	if err != nil {
		return nil, err
	}
	AddKeyQuery, err := sqlScripts.ReadFile("scripts/AddKey.sql")
	if err != nil {
		return nil, err
	}
	GetKeyQuery, err := sqlScripts.ReadFile("scripts/GetKey.sql")
	if err != nil {
		return nil, err
	}
	RemoveKeyQuery, err := sqlScripts.ReadFile("scripts/RemoveKey.sql")
	if err != nil {
		return nil, err
	}
	AddUserAgentsQuery, err := sqlScripts.ReadFile("scripts/AddUserAgent.sql")
	if err != nil {
		return nil, err
	}
	GetUserAgentsQuery, err := sqlScripts.ReadFile("scripts/GetUserAgent.sql")
	if err != nil {
		return nil, err
	}
	AddIpQuery, err := sqlScripts.ReadFile("scripts/AddIP.sql")
	if err != nil {
		return nil, err
	}
	CheckIpQuery, err := sqlScripts.ReadFile("scripts/CheckIP.sql")
	if err != nil {
		return nil, err
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
	}, nil
}
