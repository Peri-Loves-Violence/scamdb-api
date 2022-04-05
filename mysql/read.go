package mysql

import (
	"github.com/Peri-Loves-Violence/scamdb-api/sql"
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func Entry(name string, url string, user string, token string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.MySQLDB,
		User:       user,
		Token:      token,
	}
}

func ListServices(db types.ServerEntry) ([]string, error) {
	return sql.SQL(sql.MySQL, db.URL, "SELECT service FROM services")
}

func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	service_id, err := sql.SQL(sql.MySQL, db.URL, "SELECT service_id FROM services WHERE service = "+service)
	if err != nil || len(service_id) < 1 {
		return make([]string, 0), err
	}
	return sql.SQL(sql.MySQL, db.URL, "SELECT username FROM users WHERE service_id = "+service_id[0])
}

func ReadUser(user string, serv string, db types.ServerEntry) (types.UserEntry, error) {
	return sql.SQLUser(sql.MySQL, db.URL, serv, user)
}
