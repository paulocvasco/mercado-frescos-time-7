package user

type User struct {
	username string
	password string
	token    string
}

var users []User

func (u *Users) NewUser(newUser User) {
	users = append(users, newUser)
}

func (u *Users) GetUser(username string) (User, error) {
	for _, user := range users {
		if user.username == username {
			return user, nil
		}
	}
	return User{}, nil
}

type Users struct{}

type Repository interface {
	NewUser(User)
	GetUser(string) (User, error)
}

func NewRpository() Repository {
	return &Users{}
}
