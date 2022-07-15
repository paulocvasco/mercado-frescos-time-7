package mockFile

import (
	"encoding/json"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
)

type DBmock struct {
	db         interface{}
	WriteError bool
	ReadError  bool
}

func NewDatabaseMock(dbInMemory interface{}, writeError bool, readError bool) *DBmock {
	return &DBmock{
		db:         dbInMemory,
		WriteError: writeError,
		ReadError:  readError,
	}
}

func (db *DBmock) Save(model interface{}) error {
	if !db.WriteError {
		db.db = model
		return nil
	}
	return customerrors.ErrorStoreFailed
}

func (db *DBmock) Load(model interface{}) error {
	if !db.ReadError {
		dbByte, _ := json.Marshal(db.db)
		err := json.Unmarshal(dbByte, &model)
		return err
	}
	return customerrors.ErrorStoreFailed
}
