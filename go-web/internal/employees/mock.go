package employees

// type to simulete repository
type mockedRepository struct{}

func CreateMockRepository() *mockedRepository {
	return &mockedRepository{}
}

func (r *mockedRepository) Create(id int, card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error) {
	return createResponse.model, createResponse.err
}

func (r *mockedRepository) Delete(id int) error {
	return deleteResponse.err
}

func (r *mockedRepository) GetAll() ([]Employee, error) {
	return getAllResponse.modelList, getAllResponse.err
}

func (r *mockedRepository) GetByID(id int) (Employee, error) {
	return getByIdresponse.model, getByIdresponse.err
}

func (r *mockedRepository) LastID() (int, error) {
	return lastIDResponse.lastId, lastIDResponse.err
}

func (r *mockedRepository) Update(e Employee, id int) (Employee, error) {
	return updateResponse.model, updateResponse.err
}

func (r *mockedRepository) ValidationCardNumberID(card_number_id string) error {
	return validationResponse.err
}

// response to each method
var lastIDResponse struct {
	lastId int
	err    error
}

func ConfigLastID(lastID int, err error) {
	lastIDResponse.lastId = lastID
	lastIDResponse.err = err
}

var validationResponse struct {
	err error
}

func ConfigValidationCard(err error) {
	validationResponse.err = err
}

var createResponse struct {
	model Employee
	err   error
}

func ConfigCreate(model Employee, err error) {
	createResponse.model = model
	createResponse.err = err
}

var getAllResponse struct {
	modelList []Employee
	err       error
}

func ConfigureGetAll(all []Employee, err error) {
	getAllResponse.modelList = all
	getAllResponse.err = err
}

var deleteResponse struct {
	err error
}

func ConfigureDelete(err error) {
	deleteResponse.err = err
}

var getByIdresponse struct {
	model Employee
	err   error
}

func ConfigureGetByID(model Employee, err error) {
	getByIdresponse.model = model
	getByIdresponse.err = err
}

var updateResponse struct {
	model Employee
	err   error
}

func ConfigureUpdate(model Employee, err error) {
	updateResponse.model = model
	updateResponse.err = err
}
