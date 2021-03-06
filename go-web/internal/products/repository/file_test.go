package repository_test

import (
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/internal/products/mock/mockfiledb"
	"mercado-frescos-time-7/go-web/internal/products/repository"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}

	dbMock := mockfiledb.NewDatabaseMock(myDb, true, false)
	repo := repository.NewRepository(dbMock)
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
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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
	assert.Equal(t, insert1, res.Products[0])
	assert.Equal(t, insert2, res.Products[1])
	assert.Equal(t, 2, len(res.Products))
}

func TestGetAllDbFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, true, true)
	repo := repository.NewRepository(dbMock)
	res, err := repo.GetAll()
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, models.Products{}, res)
}

func TestGetByIdShouldReturnProduct(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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
	resGet1, _ := repo.GetById(insert1.Id)
	resGet2, err := repo.GetById(insert2.Id)
	assert.Nil(t, err)
	assert.Equal(t, insert1, resGet1)
	assert.Equal(t, insert2, resGet2)
}

func TestGetByInvalidIdShouldReturnErrorId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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
	_, err := repo.GetById(10)
	assert.Equal(t, customerrors.ErrorInvalidID, err)
}

func TestGetByIdDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, true)
	repo := repository.NewRepository(dbMock)
	res, err := repo.GetById(0)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
	assert.Equal(t, models.Product{}, res)
}

func TestUpdateSuccess(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
	model1 := models.Product{
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
		Id:                             1,
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
	insert, _ := repo.Insert(model1)
	repo.Update(modelEdit)
	resEdit, err := repo.GetById(insert.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, modelEdit, resEdit)
}

func TestUpdateWithInvalidId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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
	insert, _ := repo.Insert(model1)
	dbMock.WriteError = true
	err := repo.Update(insert)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestUpdateDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, true)
	repo := repository.NewRepository(dbMock)
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
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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

	insert, _ := repo.Insert(model1)
	_, errIDBefore := repo.GetById(insert.Id)
	errDelete := repo.Delete(insert.Id)
	_, errIDAfter := repo.GetById(insert.Id)
	assert.Equal(t, nil, errDelete)
	assert.Equal(t, nil, errIDBefore)
	assert.Equal(t, customerrors.ErrorInvalidID.Error(), errIDAfter.Error())
}

func TestDeleteWithInvalidId(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
	errDelete := repo.Delete(0)
	assert.Equal(t, customerrors.ErrorInvalidID, errDelete)
}

func TestDeleteDbWriteFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, false)
	repo := repository.NewRepository(dbMock)
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

	insert, _ := repo.Insert(model1)
	dbMock.WriteError = true
	err := repo.Delete(insert.Id)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}

func TestDeleteDbReadFailed(t *testing.T) {
	myDb := models.ProductMetaData{}
	dbMock := mockfiledb.NewDatabaseMock(myDb, false, true)
	repo := repository.NewRepository(dbMock)
	err := repo.Delete(0)
	assert.Equal(t, customerrors.ErrorStoreFailed.Error(), err.Error())
}
