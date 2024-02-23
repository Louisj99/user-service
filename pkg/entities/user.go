package entities

type User struct {
	ID       string
	Email    string
	Name     string
	Admin    bool
	disabled bool
}
