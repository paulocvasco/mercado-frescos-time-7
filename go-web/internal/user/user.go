package user

type User struct {
	username string
	password string
}

type Users struct{}

var users []User

func (u *Users) NewUser(newUser User) {
	users = append(users, newUser)
}
