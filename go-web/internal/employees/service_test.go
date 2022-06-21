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

			nil,
		},
	}

	for _, t := range testCases {
		repo := mock.CreateMockRepository(0, t.expectedList, Employee{}, t.expectedError)
		s := NewService(repo)
	}
}
