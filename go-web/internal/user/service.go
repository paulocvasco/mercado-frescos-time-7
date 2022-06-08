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
	s.repository.NewUser(newUser)
}

func (s *service) GetToken(user User) (string, error) {
	user, err := s.repository.GetUser(user.username)
	if err != nil {
		return "", err
	}

	return user.username, nil
}
