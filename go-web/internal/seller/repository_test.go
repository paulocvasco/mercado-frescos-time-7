package seller_test

func TestGetAll(t *testing.T) {
	type tests struct {
		name           string
		mockResponse   []models.Seller
		expectResponse []models.Seller
		expectError    error
		message        string
	}
	response := []models.Seller{
		{
			ID:           1,
			Cid:          123,
			Company_name: "Meli1",
			Address:      "Rua 1",
			Telephone:    "(11) 33387767",
		},
		{
			ID:           2,
			Cid:          1234,
			Company_name: "Meli2",
			Address:      "Rua 3",
			Telephone:    "(11) 33387768",
		},
		{
			ID:           3,
			Cid:          12356,
			Company_name: "Meli3",
			Address:      "Rua 3",
			Telephone:    "(11) 33387769",
		},
	}

	testCases := []tests{
		{"Get all Sellers", response, response, nil, "Values Differents"},
		{"GetAll return Error", nil, nil, errors.New("Error"), "Value Error Different"},
	}

	for _, value := range testCases {
		mockDB := mocks.NewDB(t)
		mockRepository := mocks.NewRepository(t)
		service := seller.NewService(mockRepository)
		mockRepository.On("GetAll").Return(value.mockResponse, value.expectError)
		resp, err := service.GetAll()
		assert.Equal(t, value.expectResponse, resp, value.name, value.message)
		assert.Equal(t, value.expectError, err, value.name, value.message)

	}

}