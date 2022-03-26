package sqlite

import (
	"database/sql"

	"github.com/Peri-Loves-Violence/scamdb-api/types"
	_ "github.com/mattn/go-sqlite3"
)

func Entry(name string, url string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.RedisDB,
		User:       "",
		Token:      "",
	}
}

func ListServices(db types.ServerEntry) ([]string, error) {
	scamDBsqlite, err := sql.Open("sqlite3", db.URL)
	if err != nil {
		return make([]string, 0), err
	}
	defer scamDBsqlite.Close()
	scamDBsqlite.Exec("SELECT service FROM ")
	return nil, nil
}

func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	return make([]string, 0), nil
}

func ReadUser(user string, serv string, db types.ServerEntry) (types.UserEntry, error) {
	return types.UserEntry{}, nil
}
