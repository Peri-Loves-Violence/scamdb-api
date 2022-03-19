package utils

import (
	"github.com/Peri-Loves-Violence/scamdb-api"
)

func NewServerEntry(name string, url string, dbType scamdb.DatabaseType, token string) scamdb.ServerEntry {
	return scamdb.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       dbType,
		Token:      token,
	}
}
