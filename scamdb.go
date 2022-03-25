package scamdb

import (
	"github.com/Peri-Loves-Violence/scamdb-api/github"
	"github.com/Peri-Loves-Violence/scamdb-api/local"
	"github.com/Peri-Loves-Violence/scamdb-api/mysql"
	"github.com/Peri-Loves-Violence/scamdb-api/redis"
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func NewDatabase(name string, typ types.DatabaseType, url string, user string, token string) types.Database {
	switch typ {
	case "github":
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

	case "mysql":
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

	case "redis":
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

	case "local":
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
	default:
		return types.Database{}
	}
}
