package scamdb

import (
	"github.com/Peri-Loves-Violence/scamdb-api/types"

	"github.com/Peri-Loves-Violence/scamdb-api/github"
	"github.com/Peri-Loves-Violence/scamdb-api/local"
	"github.com/Peri-Loves-Violence/scamdb-api/mongodb"
	"github.com/Peri-Loves-Violence/scamdb-api/mysql"
	"github.com/Peri-Loves-Violence/scamdb-api/postgresql"
	"github.com/Peri-Loves-Violence/scamdb-api/redis"
	"github.com/Peri-Loves-Violence/scamdb-api/sqlite"
)

func ScamDB(entry types.ServerEntry) types.Database {
	return NewScamDB(entry.ServerName, entry.Type, entry.URL, entry.User, entry.Token)
}

func NewScamDB(name string, typ types.DatabaseType, url string, user string, token string) types.Database {
	switch typ {
	case types.GithubDB:
		entry := github.Entry(name, url, user, token)
		return types.Database{
			Services: func() ([]string, error) {
				return github.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return github.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return github.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return github.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.SQLiteDB:
		entry := sqlite.Entry(name, url)
		return types.Database{
			Services: func() ([]string, error) {
				return sqlite.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return sqlite.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return sqlite.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return sqlite.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.MySQLDB:
		entry := mysql.Entry(name, url, user, token)
		return types.Database{
			Services: func() ([]string, error) {
				return mysql.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return mysql.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return mysql.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return mysql.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.MongoDB:
		entry := mongodb.Entry(name, url, user, token)
		return types.Database{
			Services: func() ([]string, error) {
				return mongodb.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return mongodb.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return mongodb.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return mongodb.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.RedisDB:
		entry := redis.Entry(name, url, user, token)
		return types.Database{
			Services: func() ([]string, error) {
				return redis.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return redis.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return redis.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return redis.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.LocalDB:
		entry := local.Entry(name, url)
		return types.Database{
			Services: func() ([]string, error) {
				return local.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return local.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return local.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return local.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}

	case types.PostgresDB:
		entry := postgresql.Entry(name, url, user, token)
		return types.Database{
			Services: func() ([]string, error) {
				return postgresql.ListServices(entry)
			},
			Users: func(serv string) ([]string, error) {
				return postgresql.ListUsers(serv, entry)
			},
			ReadUser: func(user string, serv string) (types.UserEntry, error) {
				return postgresql.ReadUser(user, serv, entry)
			},
			WriteUser: func(user types.UserEntry, serv string) error {
				return postgresql.WriteUser(user, serv, entry)
			},
			Entry: entry,
		}
	default:
		return types.Database{}
	}
}
