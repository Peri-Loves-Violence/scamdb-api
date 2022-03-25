package types

type UserEntry struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Profile  string `json:"profile"`
	Reason   string `json:"reason"`
}
