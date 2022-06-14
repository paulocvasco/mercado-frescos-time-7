package db

import (
	"encoding/json"
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"os"
)

type DB interface {
	Save(interface{}) error
	Load(interface{}) error
}

type database struct{}

func NewDatabase() DB {
	return &database{}
}

func (db *database) Save(model interface{}) error {
	var err error
	var file *os.File

	defer file.Close()

	path, err := getPath(model)
	if err != nil {
		return err
	}
	// check if exists a file to save data
	_, err = os.Stat(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("{}"))
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

func (db *database) Load(model interface{}) error {
	var file *os.File
	path, err := getPath(model)
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("{}"))
		if err != nil {
			return err
		}
	} else {
		file, err = os.Open(path)
		if err != nil {
			return err
		}
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

func getPath(model interface{}) (string, error) {
	switch model.(type) {
	case models.Buyer:
		return "./buyer.db", nil
	case models.Employee:
		return "./employee.db", nil
	case *models.ProductMetaData, models.ProductMetaData:
		return "./produt.db", nil
	case models.Section:
		return "./section.db", nil
	case models.Seller:
		return "./seller.db", nil
	case models.WarehouseMetaData, *models.WarehouseMetaData:
		return "./warehouse.db", nil
	default:
		return "", customerrors.ErrorInvalidDB
	}
}
