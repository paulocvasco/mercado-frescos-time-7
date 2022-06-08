package user

type Service interface {
	NewUser(User)
	GetToken(User) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) NewUser(newUser User) {
	// create token before save NewUser
	s.repository.NewUser(newUser)
}

func (s *service) GetToken(user User) (string, error) {
	user, err := s.repository.GetUser(user.Username)
	if err != nil {
		return "", err
	}

	// validate password and return token
	return user.Username, nil
}
