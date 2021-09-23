package domains

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

var Guser = User{
	Name:  "demo",
	Email: "demo@gmail.com",
	Pass:  "asdf",
}

var UserList []User
