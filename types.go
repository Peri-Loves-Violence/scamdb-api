package scamdb

// Database
type ServerEntry struct {
	ServerName string       `json:"name"`
	URL        string       `json:"url"`
	Type       DatabaseType `json:"type"`
	User       string       `json:"user"`
	Token      string       `json:"token"`
}

type DatabaseType string

const (
	GithubDB DatabaseType = "github"
	MySQLDB  DatabaseType = "mysql"
	RedisDB  DatabaseType = "redis"
	LocalDB  DatabaseType = "local"
)

// User entries
type UserEntry struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Profile  string `json:"profile"`
	Reason   string `json:"reason"`
}
