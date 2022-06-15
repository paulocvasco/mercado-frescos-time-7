package mock_DB

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/pkg/db"
)

type MockedDatabase struct {
	data interface{}
}

func (db *MockedDatabase) Save(model interface{}) error {
	dbByte, _ := json.Marshal(db.data)
	err := json.Unmarshal(dbByte, &model)
	return err
}

func (db *MockedDatabase) Load(model interface{}) error {
	dbByte, _ := json.Marshal(db.data)
	err := json.Unmarshal(dbByte, &model)
	return err
}

func NewMockedDatabase(model interface{}) db.DB {
	return &MockedDatabase{
		data: model,
	}
}