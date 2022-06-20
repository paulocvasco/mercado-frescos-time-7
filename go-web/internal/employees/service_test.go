package employees

import (
	"mercado-frescos-time-7/go-web/internal/employees/mock"
	"testing"
)

func TestGetAll(t *testing.T) {
	testCases := []struct {
		testName      string
		expectedList  []Employee
		expectedError error
	}{
		{
			"EmptyList",
			[]Employee{},
			nil,
		},
	}

	for _, t := range testCases {
		repo := mock.CreateMockRepository(0, t.expectedList, Employee{}, t.expectedError)
		s := NewService(repo)
	}
}
