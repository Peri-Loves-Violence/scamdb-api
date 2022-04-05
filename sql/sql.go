package sql

import (
	"database/sql"

	"github.com/Peri-Loves-Violence/scamdb-api/types"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type SQLType string

const (
	MySQL    SQLType = "mysql"
	SQLite   SQLType = "sqlite3"
	Postgres SQLType = "postgres"
)

func SQL(driver SQLType, url string, query string) ([]string, error) {
	scamDBsqlite, err := sql.Open(string(driver), url)
	if err != nil {
		return make([]string, 0), err
	}
	defer scamDBsqlite.Close()

	var services []string
	sqlres, err := scamDBsqlite.Query(query)
	if err != nil {
		return make([]string, 0), err
	}
	defer sqlres.Close()
	for sqlres.Next() {
		var service string
		err = sqlres.Scan(&service)
		if err != nil {
			return make([]string, 0), err
		}
		services = append(services, service)
	}
	return services, err
}

func SQLUser(driver SQLType, url string, service string, username string) (types.UserEntry, error) {
	scamDBsql, err := sql.Open(string(driver), url)
	if err != nil {
		return types.UserEntry{}, err
	}
	defer scamDBsql.Close()

	response, err := scamDBsql.Query("SELECT service_id FROM services WHERE service = $1", service)
	if err != nil {
		return types.UserEntry{}, err
	}
	defer response.Close()
	var service_id string
	err = response.Scan(&service_id)
	if err != nil {
		return types.UserEntry{}, err
	}

	sqlres, err := scamDBsql.Query("SELECT username, id, profile, reason FROM users WHERE username = $1 AND service_id = $2", username, service_id)
	if err != nil {
		return types.UserEntry{}, err
	}
	defer sqlres.Close()

	user := types.UserEntry{}
	err = sqlres.Scan(&user.Username, &user.ID, &user.Profile, &user.Reason)
	if err != nil {
		return types.UserEntry{}, err
	}
	return user, nil
}

func SQLWrite(driver SQLType, url string, user types.UserEntry, service string) error {
	scamDBsql, err := sql.Open(string(driver), url)
	if err != nil {
		return err
	}
	defer scamDBsql.Close()

	resp, err := scamDBsql.Query("SELECT service_id FROM services WHERE service = $1", service)
	if err != nil {
		return err
	}
	defer resp.Close()

	var service_id string
	resp.Scan(&service_id)

	_, err = scamDBsql.Exec("INSERT INTO users (username, id, profile, reason, service_id) VALUES ($1, $2, $3, $4, $5)", user.Username, user.ID, user.Profile, user.Reason, service_id)

	return err
}
