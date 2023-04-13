package models

type User struct {
	ID       int
	Username string
	Password string
}

// constructor
func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

// Setters
func (user *User) SetUsername(name string) {
	user.Username = name
}
func (user *User) SetPassword(password string) {
	user.Password = password
}
