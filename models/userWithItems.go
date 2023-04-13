package models

type UserWithItems struct {
	Username string
	Password string
	Items    []Item
}
