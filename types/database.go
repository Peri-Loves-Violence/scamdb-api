package types

type ServerEntry struct {
	ServerName string       `json:"name"`
	URL        string       `json:"url"`
	Type       DatabaseType `json:"type"`
	User       string       `json:"user"`
	Token      string       `json:"token"`
}

type DatabaseType string

const (
	GithubDB   DatabaseType = "github"
	MySQLDB    DatabaseType = "mysql"
	SQLiteDB   DatabaseType = "sqlite"
	PostgresDB DatabaseType = "postgres"
	LocalDB    DatabaseType = "local"
	GoogleDB   DatabaseType = "google"
)
