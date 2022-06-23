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

func TestErrorSectionNumber(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	mysec := sections.Section{
		ID:                 1,
		SectionNumber:      -1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorSectionNumber, err)
}

func TestErrorCurrentCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	mysec := sections.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    -1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorCurrentCapacity, err)
}

func TestErrorMinimumCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	mysec := sections.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    -1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorMinimumCapacity, err)
}

func TestErrorMaximumCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	mysec := sections.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    -1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorMaximumCapacity, err)
}

func TestErrorWarehouseID(t *testing.T) {
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
		WarehouseId:        -1,
		ProductTypeId:      1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorWarehouseID, err)
}

func TestErrorProductTypeID(t *testing.T) {
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
		ProductTypeId:      -1,
	}

	mbytes, _ := json.Marshal(mysec)

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorProductTypeID, err)
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
	assert.Equal(t, customErrors.ErrorConflict, err)
}

func TestCreateFailStore(t *testing.T) {
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

	repo.On("VerifySectionNumber", mock.Anything).Return(nil)
	repo.On("Store", mock.Anything).Return(sections.Section{}, customErrors.ErrorStoreFailed)
	_, err := serv.Store(mbytes)
	assert.Equal(t, customErrors.ErrorStoreFailed, err)
}

func TestCreateFailUnMarshal(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	_, err := serv.Store(nil)
	//a funcao verifica se o erro é nil porque o service retorna nil para esses caso de erro
	assert.Equal(t, nil, err)
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
	assert.Equal(t, stub, result)
}

func TestFindAllERROR(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)

	repo.On("GetAll", mock.Anything).Return(nil, errors.New("falha ao retornar dados"))
	result, errResult := serv.GetAll()
	assert.Equal(t, nil, result)
	assert.NotEqual(t, nil, errResult)
}

func TestFindByIdNE(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)

	repo.On("GetById", mock.Anything).Return(models.Section{}, customErrors.ErrorInvalidID)

	_, resultError := serv.GetById("1")
	assert.Equal(t, customErrors.ErrorInvalidID, resultError)

}

func TestFindByIdSucess(t *testing.T) {
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
	repo.On("GetById", mock.Anything).Return(mysec1, nil)

	result, resultError := serv.GetById("1")
	assert.Equal(t, mysec1, result)
	assert.Equal(t, nil, resultError)

}

func TestFindByIdERROR(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)

	_, resultError := serv.GetById("a")
	assert.Equal(t, customErrors.ErrorInvalidID, resultError)

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
	assert.Equal(t, umysec, result)
}

func TestUpdateErrorSectionNumber(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      -2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorSectionNumber, err)
}

func TestUpdateErrorCurrentCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    -1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorCurrentCapacity, err)
}

func TestUpdateErrorMinimumCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    -1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorMinimumCapacity, err)
}

func TestUpdateErrorMaximumCapacity(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    -1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorMaximumCapacity, err)
}

func TestUpdateErrorWarehouseID(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        -1,
		ProductTypeId:      1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorWarehouseID, err)
}

func TestUpdateErrorProductTypeID(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	umysec := sections.Section{
		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      -1,
	}

	//Existe uma conversão do objeto inicial models.Section para []bytes (ln. 105) e depois para sections.Section (ln. 112)

	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err:= serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorProductTypeID, err)
}

func TestUpdateNE(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)

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

	repo.On("GetById", mock.Anything).Return(models.Section{}, errors.New("erro simulado"))
	repo.On("ValidateID", mock.Anything).Return(true)
	mbytes, _ := json.Marshal(umysec)
	_, err := serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorInvalidID, err)
}

func TestUpdateErrorValidateID(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)

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
	repo.On("ValidateID", mock.Anything).Return(false)
	mbytes, _ := json.Marshal(umysec)
	_, err := serv.Update("2", mbytes)
	assert.Equal(t, customErrors.ErrorInvalidID, err)
}

func TestUpdateFailID(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	_, err := serv.Update("s", nil)
	//a funcao verifica se o erro é nil porque o service retorna nil para esses caso de erro
	assert.Equal(t, customErrors.ErrorInvalidID, err)
}

func TestDeleteNE(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	repo.On("GetById", mock.Anything).Return(models.Section{}, customErrors.ErrorInvalidID)

	_, resultError := serv.GetById("1")
	assert.Equal(t, customErrors.ErrorInvalidID, resultError)

}

func TestDEeleteOK(t *testing.T) {
	repo := mockRepository.NewRepository(t)
	serv := sections.NewService(repo)
	repo.On("GetById", mock.Anything).Return(models.Section{}, nil)
	repo.On("ValidateID", mock.Anything).Return(true)
	repo.On("Delete", mock.Anything).Return(nil)
	resultError := serv.Delete("1")
	assert.Equal(t, nil, resultError)

}
