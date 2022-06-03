package sectionRepository

type Section struct {
	ID                  int `json:"id"`
	Section_number      int `json:"section_number"`
	Current_temperature int `json:"current_temperature"`
	Minimum_temperature int `json:"minimum_temperature"`
	Current_capacity    int `json:"current_capacity"`
	Minimum_capacity    int `json:"minimum_capacity"`
	Maximim_capacity    int `json:"maximim_capacity"`
	Warehouse_id        int `json:"warehouse_id"`
	Product_type_id     int `json:"product_type_id"`
}

var sts []Section
var lastID int

type Repository interface {
	GetAll() ([]Section, error)
	GetById()
	Store(id,
		section_number,
		current_temperature,
		minimum_temperature,
		current_capacity,
		minimum_capacity,
		maximim_capacity,
		warehouse_id,
		product_type_id int) (Section, error)
	LastID() (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	repo := &repository{}
	return repo
}

func (r *repository) GetAll() ([]Section, error) {
	return sts, nil
}

func (r *repository) GetById() {
}

func (r *repository) Store(id, section_number, current_temperature, minimum_temperature,
	current_capacity, minimum_capacity, maximim_capacity, warehouse_id, product_type_id int) (Section, error) {

	section := Section{id, section_number, current_temperature, minimum_temperature,
		current_capacity, minimum_capacity, maximim_capacity, warehouse_id, product_type_id}
	sts = append(sts, section)
	lastID = section.ID

	return section, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
