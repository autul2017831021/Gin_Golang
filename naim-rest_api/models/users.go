package models

type User struct {
	Name  string
	Email string
	Pass  string
}

var Guser = User{
	Name:  "demo",
	Email: "demo@gmail.com",
	Pass:  "asdf",
}
