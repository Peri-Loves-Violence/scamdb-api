package mysql

import (
	"github.com/Peri-Loves-Violence/scamdb-api"
)

func Entry(name string, url string, user string, token string) scamdb.ServerEntry {
	return scamdb.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       scamdb.MySQLDB,
		User:       user,
		Token:      token,
	}
}
