package mock

import (
	"encoding/json"
	"log"
)

type DatabaseResponse struct {
	LoadData  string
	LoadError error
	SaveError error
}

var databaseResponse DatabaseResponse

type MockedDatabase struct{}

func (db *MockedDatabase) Save(model interface{}) error {
	return databaseResponse.SaveError
}

func (db *MockedDatabase) Load(model interface{}) error {
	err := json.Unmarshal([]byte(databaseResponse.LoadData), &model)
	if err != nil {
		log.Fatal(err)
	}
	return databaseResponse.LoadError
}

func CreateMockedDatabase(response DatabaseResponse) *MockedDatabase {
	databaseResponse.LoadData = response.LoadData
	databaseResponse.LoadError = response.LoadError
	databaseResponse.SaveError = response.SaveError

	return &MockedDatabase{}
}
