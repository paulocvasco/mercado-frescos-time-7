package seller

type Seller struct {
	ID    int  `json:"id"`
	Cid  string  `json:"cid"`
	Company_name string  `json:"company_name"`
	Address string  `json:"address"`
	Telephone  string `json:"telephone"`
}

var ps []Seller
var lastID int

type Repository interface {
	GetAll() ([]Seller, error)
	Store(id int, cid string, company_name string, address string, telephone string) (Seller, error)
	LastID() (int, error)
	}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Seller, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, cid string, company_name string, address string, telephone string) (Product, error) {
	p := Seller{id, cid, company_name, address, telephone}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}