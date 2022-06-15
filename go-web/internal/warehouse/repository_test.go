package warehouse

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/mock"
	"testing"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		testName      string
		dbResponse    mock.DatabaseResponse
		newObj        models.Warehouse
		expectedObj   models.Warehouse
		expectedError error
	}{
		{"FirstObject",
			mock.DatabaseResponse{LoadData: `{}`, LoadError: nil, SaveError: nil},
			models.Warehouse{Address: "Foo", Telephone: "xxxxxxx", MinimunCapacity: 12, MinimunTemperature: 10},
			models.Warehouse{ID: 1, Address: "Foo", Telephone: "xxxxxxx", MinimunCapacity: 12, MinimunTemperature: 10}, nil,
		},
		{"LoadFail",
			mock.DatabaseResponse{LoadData: `{}`, LoadError: customerrors.ErrorInvalidDB, SaveError: nil},
			models.Warehouse{},
			models.Warehouse{}, customerrors.ErrorInvalidDB,
		},
		{"SaveFail",
			mock.DatabaseResponse{LoadData: `{}`, LoadError: nil, SaveError: customerrors.ErrorInvalidDB},
			models.Warehouse{Address: "Foo", Telephone: "xxxxxxx", MinimunCapacity: 12, MinimunTemperature: 10},
			models.Warehouse{}, customerrors.ErrorInvalidDB,
		},
	}

	for _, v := range testCases {
		mockDB := mock.CreateMockedDatabase(v.dbResponse)
		repository := NewRepository(mockDB)

		obj, err := repository.Create(v.newObj)
		if v.expectedError != err {
			t.Errorf("Create test[%s]: error expected to be:\n%s\n\t--- but got ---\n%s\n", v.testName, v.expectedError, err)
		}

		if v.expectedObj.ID != obj.ID || v.expectedObj.Address != obj.Address || v.expectedObj.MinimunCapacity != obj.MinimunCapacity || v.expectedObj.MinimunTemperature != obj.MinimunTemperature {
			t.Errorf("Create test[%s]: object expected to be:\n%+v\n\t--- but got ---\n%+v\n", v.testName, v.expectedObj, obj)

		}
	}
}
