package sql

import (
	"fmt"
)

func PostgresURL(url string, user string, token string) string {
	return fmt.Sprintf("postgres://%s:%s@%s", user, token, url)
}

func MySQLURL(url string, user string, token string) string {
	return fmt.Sprintf("%s:%s@%s", user, token, url)
}
