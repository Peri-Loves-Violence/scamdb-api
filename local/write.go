package local

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func WriteUser(user types.UserEntry, serv string, db types.ServerEntry) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(db.URL, serv, user.Username+".json"), data, 0644)
}
