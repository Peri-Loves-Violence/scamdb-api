package postgresql

import (
	"github.com/Peri-Loves-Violence/scamdb-api/sql"
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func Entry(name string, url string, user string, token string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.PostgresDB,
		User:       user,
		Token:      token,
	}
}

func ListServices(db types.ServerEntry) ([]string, error) {
	return sql.SQL(sql.Postgres, db.URL, "SELECT name FROM services")
}

func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	return sql.SQL(sql.Postgres, db.URL, "SELECT name FROM users WHERE service = "+service)
}

func ReadUser(user string, serv string, db types.ServerEntry) (types.UserEntry, error) {
	return sql.SQLUser(sql.Postgres, db.URL, serv, user)
}
