package sections_test

import (
	"encoding/json"
	"errors"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/sections"
	"mercado-frescos-time-7/go-web/internal/sections/mock/mockRepository"
	customErrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestCreateOK(t *testing.T) {
	repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)
    mysec := sections.Section{
        ID:                 1,
        SectionNumber:      1,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    1,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }

	mbytes, _ := json.Marshal(mysec)

    repo.On("Store", mock.Anything).Return(mysec, nil)
	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
    result, _ := serv.Store(mbytes)
    assert.Equal(t, mysec, result)
}

func TestCreateFailConflict(t *testing.T) {
    repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)
    mysec1 := models.Section{
        ID:                 1,
        SectionNumber:      1,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    1,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }

	mbytes, _ := json.Marshal(mysec1)

    repo.On("VerifySectionNumber", mock.Anything).Return(customErrors.ErrorConflict)
    _, err := serv.Store(mbytes)
    assert.Equal(t, customErrors.ErrorConflict, err,)
}

func TestFindAll(t *testing.T) {
    repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)
    stub := []models.Section{   
            {
                ID:                 1,
                SectionNumber:      1,
                CurrentTemperature: 1,
                MinimumTemperature: 1,
                CurrentCapacity:    -1,
                MinimumCapacity:    1,
                MaximumCapacity:    1,
                WarehouseId:        1,
                ProductTypeId:      1,
            },
            {
                ID:                 2,
                SectionNumber:      1,
                CurrentTemperature: 1,
                MinimumTemperature: 1,
                CurrentCapacity:    -1,
                MinimumCapacity:    1,
                MaximumCapacity:    1,
                WarehouseId:        1,
                ProductTypeId:      1,
            },
    }

    repo.On("GetAll", mock.Anything).Return(stub, nil)
    result, _ := serv.GetAll()
    assert.Equal(t, stub, result,)
}

func TestFindByIdNE(t *testing.T){
    repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)

    repo.On("GetById", mock.Anything).Return(models.Section{}, customErrors.ErrorInvalidID)
    
    _, resultError := serv.GetById("1")
    assert.Equal(t,customErrors.ErrorInvalidID, resultError,)

}

func TestUpdate(t *testing.T) {
    repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)
    mysec := models.Section{
        ID:                 1,
        SectionNumber:      1,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    1,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }
    umysec := sections.Section{
        ID:                 1,
        SectionNumber:      2,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    1,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

    repo.On("Update", mock.Anything, mock.Anything).Return(nil)
    repo.On("GetById", mock.Anything).Return(mysec, nil)
    repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
    result, _ := serv.Update("2", mbytes)
    assert.Equal(t, umysec, result,)
}

func TestUpdateNE(t *testing.T) {
    repo := mockRepository.NewRepository(t)
    serv := sections.NewService(repo)
    mysec := models.Section{
        ID:                 1,
        SectionNumber:      1,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    1,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }
    umysec := models.Section{
        ID:                 1,
        SectionNumber:      0,
        CurrentTemperature: 1,
        MinimumTemperature: 1,
        CurrentCapacity:    0,
        MinimumCapacity:    1,
        MaximumCapacity:    1,
        WarehouseId:        1,
        ProductTypeId:      1,
    }

    //repo.On("Update", mock.Anything, mock.Anything).Return(umysec, nil)
    repo.On("GetById", mock.Anything).Return(mysec, errors.New("esperado"))
    repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
    _, err := serv.Update("2", mbytes)
    assert.Equal(t, customErrors.ErrorInvalidID, err,)
}

