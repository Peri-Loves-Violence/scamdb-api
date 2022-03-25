package github

import (
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func Entry(name string, url string, user string, token string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.GithubDB,
		User:       user,
		Token:      token,
	}
}

func ListServices(db types.ServerEntry) ([]string, error) {
	return nil, nil
}

func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	return make([]string, 0), nil
}

func ReadUser(user string, serv string, db types.ServerEntry) (types.UserEntry, error) {
	return types.UserEntry{}, nil
}
