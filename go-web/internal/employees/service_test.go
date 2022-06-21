package employees

import (
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
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

	for _, v := range testCases {
		repo := CreateMockRepository()
		ConfigGetAll(v.expectedList, v.expectedError)
		s := NewService(repo)
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

	for _, v := range testCases {
		repo := CreateMockRepository()
		ConfigDelete(v.respnseError)
		s := NewService(repo)
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

	for _, v := range testCases {
		repo := CreateMockRepository()
		ConfigGetByID(v.responseModel, v.responseError)
		s := NewService(repo)
		model, err := s.GetByID(v.requestID)

		if v.expectedError != err {
			t.Errorf("Delete test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedModel != model {
			t.Errorf("GetByID test[%s]: model expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedModel, model)
		}
	}
}

			nil,
		},
	}

	for _, t := range testCases {
		repo := mock.CreateMockRepository(0, t.expectedList, Employee{}, t.expectedError)
		s := NewService(repo)
	}
}
