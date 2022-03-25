package types

type Database struct {
	Entry     ServerEntry
	Services  func() ([]string, error)
	Users     func(string) ([]string, error)
	ReadUser  func(string, string) (UserEntry, error)
	WriteUser func(UserEntry, string) error
}
