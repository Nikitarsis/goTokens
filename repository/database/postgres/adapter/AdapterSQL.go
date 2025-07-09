package postgres

import (
	"database/sql"
	"net"

	co "github.com/Nikitarsis/goTokens/common"
)

type adapterSQL struct {
	Exec                func(query string, args ...interface{}) (sql.Result, error)
	Query               func(query string, args ...interface{}) (*sql.Rows, error)
	savePorts 					bool
	CreateTablesQuery   string
	AddKeyQuery         string
	GetKeyQuery         string
	RemoveKeyQuery      string
	AddUserAgentsQuery  string
	GetUserAgentsQuery  string
	AddIpQuery          string
	CheckIpQuery        string
}

func (a *adapterSQL) CreateTablesIFNotExists() {
	a.Exec(a.CreateTablesQuery)
}

func (a *adapterSQL) AddKey(kid co.UUID, key co.Key) {
	a.Exec(a.AddKeyQuery, kid.ToString(), key.ToString())
}

func (a *adapterSQL) RemoveKey(kid co.UUID) {
	a.Exec(a.RemoveKeyQuery, kid.ToString())
}

func (a *adapterSQL) GetKey(kid co.UUID) (co.Key, bool) {
	rows, err := a.Query(a.GetKeyQuery, kid.ToString())
	if err != nil {
		return co.Key{}, false
	}
	defer rows.Close()

	var key string
	if rows.Next() {
		err := rows.Scan(&key)
		if err != nil {
			return co.Key{}, false
		}
		ret, err := co.CreateKeyFromString(kid, key)
		if err != nil {
			return co.Key{}, false
		}
		return ret, true
	}
	return co.Key{}, false
}

func (a *adapterSQL) AddUserAgent(kid co.UUID, userAgent string) {
	a.Exec(a.AddUserAgentsQuery, kid.ToString(), userAgent)
}

func (a *adapterSQL) GetUserAgent(kid co.UUID) (string, error) {
	rows, err := a.Query(a.GetUserAgentsQuery, kid.ToString())
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var userAgent string
	if rows.Next() {
		err := rows.Scan(&userAgent)
		if err != nil {
			return "", err
		}
	}
	return userAgent, nil
}

func (a *adapterSQL) AddIp(ip co.DataIP) {
	a.Exec(a.AddIpQuery, ip.KeyId.ToString(), ip.UserId.ToString(), ip.IP.String(), ip.Port)
}

func (a *adapterSQL) CheckIp(ip co.DataIP) bool {
	rows, err := a.Query(a.CheckIpQuery, ip.KeyId.ToString(), ip.UserId.ToString(), ip.IP.String(), ip.Port)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var port uint16
		var addr net.IP
		if err := rows.Scan(&port, &addr); err != nil {
			continue
		}
		if addr.Equal(ip.IP) {
			if a.savePorts {
				return port == ip.Port
			} else {
				return true
			}
		}
	}
	return false
}
