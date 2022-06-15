package Seller_test

import (
	"fmt"
	"mercado-frescos-time-7/go-web/internal/Seller"
	"mercado-frescos-time-7/go-web/internal/models"
	"mercado-frescos-time-7/go-web/pkg/db/mock/mock_DB"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T)()  {

	testCases := []struct {
		testName      string
		dbResponse    interface{}
		newObj        models.Seller
		expectedObj   []models.Seller
		expectedError error
	}{
		{"CheckList",
			models.Sellers{Seller: []models.Seller{{ID: 1, Cid: 11, Company_name: "apple", Address: "Rua G", Telephone: "998988978"},{ID: 2, Cid: 22, Company_name: "apple", Address: "Rua G", Telephone: "998988978"}}, LastID: 11},
			models.Seller{},
			[]models.Seller{{ID: 1, Cid: 11, Company_name: "apple", Address: "Rua G", Telephone: "998988978"},{ID: 2, Cid: 22, Company_name: "apple", Address: "Rua G", Telephone: "998988978"}},
			nil,
		},
	}

	for _, v := range testCases {
		bd := mock_DB.NewMockedDatabase(v.dbResponse)
		repo := Seller.NewRepository(bd)
		result, err := repo.GetAll()
		if err != nil {
			fmt.Println(err)	
			}
		assert.Equal(t, v.expectedObj, result, "resultados devem ser iguais") 
	}
	/*
	stubBD := []models.Seller{{ID: 12, Cid: 10, Company_name: "apple", Address: "Rua G", Telephone: "998988978"}}	
	mmock := models.Sellers{
		Seller: stubBD,
		LastID: 12,
	}
	bd := mock_DB.NewMockedDatabase(mmock)
	repo := Seller.NewRepository(bd)
	result, _ := repo.GetAll()
	assert.Equal(t, stubBD, result, "resultados devem ser iguais") 
	*/
}


func TestStore(t *testing.T)()  {

	testCases := []struct {
		testName      string
		dbResponse    interface{}
		newObj        models.Seller
		expectedObj   models.Seller
		expectedError error
	}{
		{"FirstTest",
			models.Sellers{Seller:[]models.Seller{}, LastID: 11},
			models.Seller{ID: 12, Cid: 10, Company_name: "apple", Address: "Rua G", Telephone: "998988978"},
			models.Seller{ID: 12, Cid: 10, Company_name: "apple", Address: "Rua G", Telephone: "998988978"},
			nil,
		},
	}

	for _, v := range testCases {
		bd := mock_DB.NewMockedDatabase(v.dbResponse)
		repo := Seller.NewRepository(bd)
		result, err := repo.Store(v.newObj.ID, v.newObj.Cid,v.newObj.Company_name,v.newObj.Address,v.newObj.Telephone)
		if err != nil {
			fmt.Println(err)	
			}
		assert.Equal(t, v.expectedObj, result, "resultados devem ser iguais") 
	}
	/*
	stubBD := models.Seller{ID: 12, Cid: 10, Company_name: "apple", Address: "Rua G", Telephone: "998988978"}	
	mmock := models.Sellers{
		Seller: []models.Seller{},
		LastID: 11,
	}
	bd := mock_DB.NewMockedDatabase(mmock)
	repo := Seller.NewRepository(bd)
	result, err := repo.Store(12, 10,"apple","Rua G","998988978")
	if err != nil {
	fmt.Println(err)	
	}
	assert.Equal(t, stubBD, result, "resultados devem ser iguais") 
	*/
}