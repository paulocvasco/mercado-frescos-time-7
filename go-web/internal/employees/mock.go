package employees

// Service and Repository has the same return value
// to simplicity use the same response
var lastIDResponse struct {
	lastId int
	err    error
}

var validationResponse struct {
	err error
}

var createResponse struct {
	model Employee
	err   error
}

var getAllResponse struct {
	modelList []Employee
	err       error
}

var deleteResponse struct {
	err error
}

var getByIdresponse struct {
	model Employee
	err   error
}

var updateResponse struct {
	model Employee
	err   error
}

func ConfigLastID(lastID int, err error) {
	lastIDResponse.lastId = lastID
	lastIDResponse.err = err
}

func ConfigValidationCard(err error) {
	validationResponse.err = err
}

func ConfigCreate(err error) {
	createResponse.err = err
}

func ConfigGetAll(all []Employee, err error) {
	getAllResponse.modelList = all
	getAllResponse.err = err
}

func ConfigDelete(err error) {
	deleteResponse.err = err
}

func ConfigGetByID(model Employee, err error) {
	getByIdresponse.model = model
	getByIdresponse.err = err
}

func ConfigUpdate(model Employee, err error) {
	updateResponse.model = model
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
	return getAllResponse.modelList, getAllResponse.err
}

func (s *mockedService) GetByID(id int) (Employee, error) {
	return getByIdresponse.model, getByIdresponse.err
}

func (s *mockedService) Create(card_number_id string, first_name string, last_name string, warehouse_id int) (Employee, error) {
	return Employee{CardNumberId: card_number_id, FirstName: first_name, LastName: last_name, WareHouseId: warehouse_id}, createResponse.err
}

func (s *mockedService) Update(e RequestPatch, id int) (Employee, error) {
	return updateResponse.model, updateResponse.err
}

func (s *mockedService) Delete(id int) error {
	return deleteResponse.err
}

///////////////////////////////////////////////////////////////////////////////////////
//                                   MOCK REPOSITORY                                 //
///////////////////////////////////////////////////////////////////////////////////////

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
