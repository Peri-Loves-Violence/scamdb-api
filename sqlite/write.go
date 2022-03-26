package sqlite

import (
	"github.com/Peri-Loves-Violence/scamdb-api/sql"
	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func WriteUser(user types.UserEntry, serv string, db types.ServerEntry) error {
	return sql.SQLWrite(sql.SQLite, db.URL, user, serv)
}
