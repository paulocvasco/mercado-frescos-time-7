package db

import (
	"encoding/json"
	"os"
)

type DB interface {
	Save(string, interface{}) error
	Load(string, interface{}) error
}

type database struct{}

func NewDatabase() DB {
	return &database{}
}

func (db *database) Save(path string, model interface{}) error {
	var err error
	var file *os.File

	defer file.Close()
	// check if exists a file to save data
	_, err = os.Stat(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
	} else {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}
	}

	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	// save data
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (db *database) Load(path string, model interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return err
	}

	return nil
}
