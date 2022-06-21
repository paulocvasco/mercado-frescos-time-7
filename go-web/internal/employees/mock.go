package employees

// type to simulete repository
type mockedRepository struct{}

func CreateMockRepository() *mockedRepository {
	return &mockedRepository{}
}

func (r *mockedRepository) Create(id int, card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error) {
	return Employee{ID: id, CardNumberId: card_number_id, FirstName: first_name, LastName: last_name, WareHouseId: warehouse_id}, createResponse.err
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
	return e, updateResponse.err
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

func ConfigCreate(err error) {
	createResponse.err = err
}

var getAllResponse struct {
	modelList []Employee
	err       error
}

func ConfigGetAll(all []Employee, err error) {
	getAllResponse.modelList = all
	getAllResponse.err = err
}

var deleteResponse struct {
	err error
}

func ConfigDelete(err error) {
	deleteResponse.err = err
}

var getByIdresponse struct {
	model Employee
	err   error
}

func ConfigGetByID(model Employee, err error) {
	getByIdresponse.model = model
	getByIdresponse.err = err
}

var updateResponse struct {
	err error
}

func ConfigUpdate(err error) {
	updateResponse.err = err
}

///////////////////////////////////////////////////////////////////////////////////////
//                                   MOCK SERVICE                                    //
///////////////////////////////////////////////////////////////////////////////////////

type mockedService struct{}

func NewMockedService() *mockedService {
	return &mockedService{}
}

func (s *mockedService) GetAll() ([]Employee, error) {
	return nil, nil
}

func (s *mockedService) GetByID(id int) (Employee, error) {
	return Employee{}, nil
}

func (s *mockedService) Create(card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error) {
	return Employee{}, nil
}

func (s *mockedService) Update(e RequestPatch, id int) (Employee, error) {
	return Employee{}, nil
}

func (s *mockedService) Delete(id int) error {
	return nil
}
