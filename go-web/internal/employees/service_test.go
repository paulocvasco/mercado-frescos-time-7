package employees

import (
	customerrors "github.com/paulocvasco/mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"
)

func TestGetAll(t *testing.T) {
	testCases := []struct {
		testName      string
		responseList  []Employee
		responseError error
		expectedList  []Employee
		expectedError error
	}{
		{
			"EmptyList", []Employee{}, nil,
			[]Employee{}, nil,
		},
		{
			"CheckList",
			[]Employee{
				{ID: 1, CardNumberId: "1", FirstName: "Foo", LastName: "Bar", WareHouseId: 1},
				{ID: 2, CardNumberId: "23", FirstName: "Fbar", LastName: "Foo", WareHouseId: 45},
			}, nil,
			[]Employee{
				{ID: 1, CardNumberId: "1", FirstName: "Foo", LastName: "Bar", WareHouseId: 1},
				{ID: 2, CardNumberId: "23", FirstName: "Fbar", LastName: "Foo", WareHouseId: 45},
			}, nil,
		},
		{
			"ErrorGetList",
			[]Employee{}, customerrors.ErrorItemNotFound,
			[]Employee{}, customerrors.ErrorItemNotFound,
		},
	}

	repo := CreateMockRepository()
	s := NewService(repo)

	for _, v := range testCases {
		ConfigGetAll(v.expectedList, v.expectedError)
		list, err := s.GetAll()

		if v.expectedError != err {
			t.Errorf("GetAll test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
			continue
		}

		for i, item := range list {
			if len(v.expectedList) != len(list) || item != v.expectedList[i] {
				t.Errorf("GetAll test[%s]: list expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedList, list)
			}
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		testName      string
		requestID     int
		respnseError  error
		expectedError error
	}{
		{
			"InvalidID",
			1,
			customerrors.ErrorInvalidID,
			customerrors.ErrorInvalidID,
		},
		{
			"Success",
			3,
			nil,
			nil,
		},
	}

	repo := CreateMockRepository()
	s := NewService(repo)

	for _, v := range testCases {
		ConfigDelete(v.respnseError)
		err := s.Delete(v.requestID)

		if v.expectedError != err {
			t.Errorf("Delete test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		testName      string
		requestID     int
		responseModel Employee
		responseError error
		expectedModel Employee
		expectedError error
	}{
		{
			"InvalidID", 1,
			Employee{}, customerrors.ErrorInvalidID,
			Employee{}, customerrors.ErrorInvalidID,
		},
		{
			"Success", 1,
			Employee{ID: 1, CardNumberId: "23", FirstName: "Foo", LastName: "Bar", WareHouseId: 3},
			nil,
			Employee{ID: 1, CardNumberId: "23", FirstName: "Foo", LastName: "Bar", WareHouseId: 3},
			nil,
		},
	}

	repo := CreateMockRepository()
	s := NewService(repo)

	for _, v := range testCases {
		ConfigGetByID(v.responseModel, v.responseError)
		model, err := s.GetByID(v.requestID)

		if v.expectedError != err {
			t.Errorf("Delete test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != model {
			t.Errorf("GetByID test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, model)
		}
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		testName        string
		cardID          string
		firstName       string
		lastName        string
		warehouseID     int
		validationError error
		lastId          int
		lastIdError     error
		createError     error
		expectedModel   Employee
		expectedError   error
	}{
		{
			"InvalidCardNumber", "0", "Foo", "Bar", 1, customerrors.ErrorCardIdAlreadyExists,
			1, nil, nil, Employee{}, customerrors.ErrorCardIdAlreadyExists,
		},
		{
			"InvalidCardNumber", "0", "Foo", "Bar", 1, nil,
			0, customerrors.ErrorInvalidID, nil, Employee{}, customerrors.ErrorInvalidID,
		},
		{
			"CreateFail", "0", "Foo", "Bar", 1, nil,
			0, nil, customerrors.ErrorInvalidDB, Employee{}, customerrors.ErrorInvalidDB,
		},
		{
			"Success", "0", "Foo", "Bar", 1, nil,
			1, nil, nil, Employee{ID: 2, CardNumberId: "0", FirstName: "Foo", LastName: "Bar", WareHouseId: 1}, nil,
		},
	}

	repo := CreateMockRepository()
	s := NewService(repo)

	for _, v := range testCases {
		ConfigValidationCard(v.validationError)
		ConfigLastID(v.lastId, v.lastIdError)
		ConfigCreate(v.createError)

		model, err := s.Create(v.cardID, v.firstName, v.lastName, v.warehouseID)
		if v.expectedError != err {
			t.Errorf("Create test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != model {
			t.Errorf("Create test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, model)
		}
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		testName           string
		updatedModel       RequestPatch
		id                 int
		validationError    error
		getIdModelResponse Employee
		getIdError         error
		updateError        error
		expectedModel      Employee
		expectedError      error
	}{
		{
			"FailCardID", RequestPatch{}, 1,
			customerrors.ErrorConflict,
			Employee{}, nil,
			nil,
			Employee{}, customerrors.ErrorConflict,
		},
		{
			"FailGetID", RequestPatch{}, 1,
			nil,
			Employee{}, customerrors.ErrorInvalidID,
			nil,
			Employee{}, customerrors.ErrorInvalidID,
		},
		{
			"ItemNotFound", RequestPatch{}, 1,
			nil,
			Employee{}, nil,
			customerrors.ErrorItemNotFound,
			Employee{}, customerrors.ErrorItemNotFound,
		},
		{
			"Success", RequestPatch{CardNumberId: "12", FirstName: "Bar", LastName: "Foo", WareHouseId: 10}, 1,
			nil,
			Employee{ID: 1, CardNumberId: "1", FirstName: "Foo", LastName: "Bar", WareHouseId: 1}, nil,
			nil,
			Employee{ID: 1, CardNumberId: "12", FirstName: "Bar", LastName: "Foo", WareHouseId: 10}, nil,
		},
	}

	repo := CreateMockRepository()
	s := NewService(repo)

	for _, v := range testCases {
		ConfigValidationCard(v.validationError)
		ConfigGetByID(v.getIdModelResponse, v.getIdError)
		ConfigUpdate(Employee{}, v.updateError)

		model, err := s.Update(v.updatedModel, v.id)
		if v.expectedError != err {
			t.Errorf("Update test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != model {
			t.Errorf("Update test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, model)
		}
	}
}
