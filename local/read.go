package local

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Peri-Loves-Violence/scamdb-api"
)

func Entry(name string, url string) scamdb.ServerEntry {
	return scamdb.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       scamdb.LocalDB,
		User:       "",
		Token:      "",
	}
}

// Lists all services in the database
func ListServices(db scamdb.ServerEntry) ([]string, error) {
	return listDir(db.URL)
}

// Lists all user entries in the database for the current service
func ListUsers(service string, db scamdb.ServerEntry) ([]string, error) {
	return listFiles(filepath.Join(db.URL, service))
}

// Returns the user object for given database with service and entry id/name
func ReadUser(username string, service string, db scamdb.ServerEntry) (scamdb.UserEntry, error) {
	userFile, err := os.ReadFile(filepath.Join(db.URL, service, username))
	if err != nil {
		return scamdb.UserEntry{}, err
	}
	var user scamdb.UserEntry
	err = json.Unmarshal(userFile, &user)
	return user, err
}
