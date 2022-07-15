package mockFile

import "encoding/json"

type DBmock struct {
	resultsMock interface{}
	WriteError  error
	ReadError   error
}

func NewDatabaseMock(resultMock interface{}, writeError, readError error) DBmock {
	return DBmock{
		resultsMock: resultMock,
		WriteError:  writeError,
		ReadError:   readError,
	}
}
func (db *DBmock) ConfigResult(buyers interface{}) {
	db.resultsMock = buyers
}

func (db *DBmock) Save(model interface{}) error {
	return db.WriteError
}

func (db *DBmock) Load(model interface{}) error {
	changeJson, _ := json.Marshal(db.resultsMock)
	json.Unmarshal(changeJson, &model)

	return db.ReadError
}
