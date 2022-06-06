package employees

type Service interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (Employee, error)
	Create(id int, card_number_id int, first_name string, last_name string, warehouse_id int) (Employee, error)
	Update(e Employee, id int) (Employee, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

// Create implements Service
func (s *service) Create(id int, card_number_id int, first_name string, last_name string, warehouse_id int) (Employee, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Employee{}, err
	}
	lastID++

	employee, err := s.repository.Create(lastID, card_number_id, first_name, last_name, warehouse_id)

	if err != nil {
		return Employee{}, err
	}

	return employee, nil

}

func (s *service) GetAll() ([]Employee, error) {
	employees, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return employees, nil

}

// Delete implements Service
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}

// GetByID implements Service
func (s *service) GetByID(id int) (Employee, error) {
	employees, err := s.repository.GetByID(id)

	if err != nil {
		return Employee{}, err
	}

	return employees, nil
}

// Update implements Service
func (s *service) Update(e Employee, id int) (Employee, error) {
	employees, err := s.repository.Update(e, id)

	if err != nil {
		return Employee{}, err
	}

	return employees, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
