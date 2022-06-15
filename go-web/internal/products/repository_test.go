package products_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/products"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db/mock/mock_DB"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, true, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "mod",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	_, err := repo.Insert(model1)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestGetAllSuccess(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	model2 := models.Product{
		Id:                             1,
		Description:                    "test 1",
		ExpirationRate:                 11,
		FreezingRate:                   21,
		Height:                         61.4,
		Length:                         41.5,
		NetWeight:                      31.4,
		ProductCode:                    "ssd1",
		RecommendedFreezingTemperature: 11.3,
		Width:                          11.2,
		ProductTypeId:                  21,
		SellerId:                       21,
	}
	insert1, _ := repo.Insert(model1)
	insert2, _ := repo.Insert(model2)
	res, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, insert1, res[0])
	assert.Equal(t, insert2, res[1])
	assert.Equal(t, 2, len(res))
}

func TestGetAllDbFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, true, true)
	repo := products.NewRepository(dbMock)
	res, err := repo.GetAll()
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, []models.Product{}, res)
}

func TestGetByIdShouldReturnProduct(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	model2 := models.Product{
		Id:                             1,
		Description:                    "test 1",
		ExpirationRate:                 11,
		FreezingRate:                   21,
		Height:                         61.4,
		Length:                         41.5,
		NetWeight:                      31.4,
		ProductCode:                    "ssd1",
		RecommendedFreezingTemperature: 11.3,
		Width:                          11.2,
		ProductTypeId:                  21,
		SellerId:                       21,
	}
	insert1, _ := repo.Insert(model1)
	insert2, _ := repo.Insert(model2)
	resGet1, _ := repo.GetById(0)
	resGet2, err := repo.GetById(1)
	assert.Nil(t, err)
	assert.Equal(t, insert1, resGet1)
	assert.Equal(t, insert2, resGet2)
}

func TestGetByInvalidIdShouldReturnErrorId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	repo.Insert(model1)
	_, err := repo.GetById(1)
	assert.Equal(t, customerrors.ErrorInvalidID, err)
}

func TestGetByIdDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, true)
	repo := products.NewRepository(dbMock)
	res, err := repo.GetById(0)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, models.Product{}, res)
}

func TestUpdateSuccess(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	modelEdit := models.Product{
		Id:                             0,
		Description:                    "Editado",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	repo.Insert(model1)
	repo.Update(modelEdit)
	resEdit, err := repo.GetById(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, modelEdit, resEdit)
}

func TestUpdateWithInvalidId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	modelEdit := models.Product{
		Id:                             2,
		Description:                    "Editado",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	repo.Insert(model1)
	err := repo.Update(modelEdit)
	assert.Equal(t, customerrors.ErrorInvalidID.Error(), err.Error())
}

func TestUpdateDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "mod",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	repo.Insert(model1)
	dbMock.WriteError = true
	err := repo.Update(model1)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestUpdateDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, true)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "mod",
		ExpirationRate:                 10,
		FreezingRate:                   20,
		Height:                         6.40,
		Length:                         4.50,
		NetWeight:                      3.40,
		ProductCode:                    "ssd-Editado",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}
	repo.Insert(model1)
	err := repo.Update(model1)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestDeleteSuccess(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	repo.Insert(model1)
	_, errIDBefore := repo.GetById(0)
	errDelete := repo.Delete(0)
	_, errIDAfter := repo.GetById(0)
	assert.Equal(t, nil, errDelete)
	assert.Equal(t, nil, errIDBefore)
	assert.Equal(t, customerrors.ErrorInvalidID.Error(), errIDAfter.Error())
}

func TestDeleteWithInvalidId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	errDelete := repo.Delete(0)
	assert.Equal(t, customerrors.ErrorInvalidID, errDelete)
}

func TestDeleteDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	model1 := models.Product{
		Id:                             0,
		Description:                    "test",
		ExpirationRate:                 1,
		FreezingRate:                   2,
		Height:                         6.4,
		Length:                         4.5,
		NetWeight:                      3.4,
		ProductCode:                    "ssd",
		RecommendedFreezingTemperature: 1.3,
		Width:                          1.2,
		ProductTypeId:                  2,
		SellerId:                       2,
	}

	repo.Insert(model1)
	dbMock.WriteError = true
	err := repo.Delete(0)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestDeleteDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, true)
	repo := products.NewRepository(dbMock)
	err := repo.Delete(0)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestLastId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, false)
	repo := products.NewRepository(dbMock)
	repo.LastId()
	repo.LastId()
	id, err := repo.LastId()
	assert.Nil(t, err)
	assert.Equal(t, 3, id)
}

func TestLastIdDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, true, false)
	repo := products.NewRepository(dbMock)
	res, err := repo.LastId()
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, 0, res)
}

func TestLastIdDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mock_DB.NewDatabaseMock(myDb, false, true)
	repo := products.NewRepository(dbMock)
	res, err := repo.LastId()
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, 0, res)
}
