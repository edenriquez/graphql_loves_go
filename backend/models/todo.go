package models

type Todo struct {
	ID   string
	Text string
	Done bool
	User User
}
