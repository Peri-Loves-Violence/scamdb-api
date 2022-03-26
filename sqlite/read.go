package sqlite

import (
	"github.com/Peri-Loves-Violence/scamdb-api/sql"
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func Entry(name string, url string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.SQLiteDB,
		User:       "",
		Token:      "",
	}
}

func ListServices(db types.ServerEntry) ([]string, error) {
	return sql.SQL(sql.SQLite, db.URL, "SELECT name FROM services")
}

func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	return sql.SQL(sql.SQLite, db.URL, "SELECT name FROM users WHERE service = "+service)
}

func ReadUser(user string, serv string, db types.ServerEntry) (types.UserEntry, error) {
	return sql.SQLUser(sql.SQLite, db.URL, serv, user)
}
