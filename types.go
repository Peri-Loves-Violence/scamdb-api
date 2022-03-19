package scamdb

type ServerEntry struct {
	ServerName string       `json:"name"`
	URL        string       `json:"url"`
	Type       DatabaseType `json:"type"`
	Token      string       `json:"token"`
}

type DatabaseType string

const (
	GithubDB   DatabaseType = "github"
	MySQLDB    DatabaseType = "mysql"
	RedisDB    DatabaseType = "redis"
	LocalDB    DatabaseType = "local"
	GitcloneDB DatabaseType = "gitclone"
)
