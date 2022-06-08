package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

var users []User

func (u *Users) NewUser(newUser User) {
	users = append(users, newUser)
}

func (u *Users) GetUser(username string) (User, error) {
	for _, user := range users {
		if user.Username == username {
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
