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

type mockedDatabase struct{}

func (db *mockedDatabase) Save(model interface{}) error {
	return databaseResponse.SaveError
}

func (db *mockedDatabase) Load(model interface{}) error {
	err := json.Unmarshal([]byte(databaseResponse.LoadData), &model)
	if err != nil {
		log.Fatal(err)
	}
	return databaseResponse.LoadError
}

type MockedDatabase interface {
	Save(interface{}) error
	Load(interface{}) error
}

func CreateMockedDatabase(response DatabaseResponse) MockedDatabase {
	databaseResponse.LoadData = response.LoadData
	databaseResponse.LoadError = response.LoadError
	databaseResponse.SaveError = response.SaveError

	return &mockedDatabase{}
}
