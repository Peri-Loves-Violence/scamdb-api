package local

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func Entry(name string, url string) types.ServerEntry {
	return types.ServerEntry{
		ServerName: name,
		URL:        url,
		Type:       types.LocalDB,
		User:       "",
		Token:      "",
	}
}

// Lists all services in the database
func ListServices(db types.ServerEntry) ([]string, error) {
	files, err := ioutil.ReadDir(db.URL)
	if err != nil {
		return make([]string, 0), err
	}
	var names []string
	for _, f := range files {
		if f.IsDir() {
			names = append(names, f.Name())
		}
	}
	return names, nil
}

// Lists all user entries in the database for the current service
func ListUsers(service string, db types.ServerEntry) ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(db.URL, service))
	if err != nil {
		return make([]string, 0), err
	}
	var names []string
	for _, f := range files {
		if !f.IsDir() {
			names = append(names, f.Name())
		}
	}
	return names, nil
}

// Returns the user object for given database with service and entry id/name
func ReadUser(username string, service string, db types.ServerEntry) (types.UserEntry, error) {
	userFile, err := os.ReadFile(filepath.Join(db.URL, service, username))
	if err != nil {
		return types.UserEntry{}, err
	}
	var user types.UserEntry
	err = json.Unmarshal(userFile, &user)
	return user, err
}
