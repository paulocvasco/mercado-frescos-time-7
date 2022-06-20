package mock

import "mercado-frescos-time-7/go-web/internal/models"

type Employee models.Employee

type repository struct{}

var repositoryResponse struct {
	lastID        int
	employees     []Employee
	returnedModel Employee
	error         error
}

func CreateMockRepository(mockedId int, mockedEmployeeList []Employee, mockedModel Employee, mockedError error) repository {
	repositoryResponse.lastID = mockedId
	repositoryResponse.employees = mockedEmployeeList
	repositoryResponse.returnedModel = mockedModel
	repositoryResponse.error = mockedError

	return repository{}
}

func (r *repository) Create(id int, cardNumber string, firstName string, lastName string, warehouseID int) (Employee, error) {
	return repositoryResponse.returnedModel, repositoryResponse.error
}

func (r *repository) GetAll() ([]Employee, error) {
	return repositoryResponse.employees, repositoryResponse.error
}

func (r *repository) Delete(id int) error {
	return repositoryResponse.error
}

func (r *repository) GetByID(id int) (Employee, error) {
	return repositoryResponse.returnedModel, repositoryResponse.error
}

func (r *repository) LastID() (int, error) {
	return repositoryResponse.lastID, repositoryResponse.error
}

func (r *repository) Update(employee Employee, id int) (Employee, error) {
	return repositoryResponse.returnedModel, repositoryResponse.error
}
