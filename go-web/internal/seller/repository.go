package seller

import (
	"mercado-frescos-time-7/go-web/internal/models"
	customerrors "mercado-frescos-time-7/go-web/pkg/custom_errors"
	"mercado-frescos-time-7/go-web/pkg/db"
)

type Seller models.Seller
type Sellers models.Sellers

var storage models.Sellers

type Repository interface {
	GetAll() ([]models.Seller, error)
	GetId(indice int) (models.Seller, error)
	CheckCid(cid int) (models.Seller, error)
	Update(s Seller, id int) (models.Seller, error)
	Delete(id int) error
	Store(sel models.Seller) (models.Seller, error)
	LastID() (int, error)
}

type repository struct {
	database db.DB
}

func NewRepository(database db.DB) Repository {
	return &repository{
		database: database,
	}
}

func (r *repository) GetAll() ([]models.Seller, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return []models.Seller{}, err
	}
	return storage.Seller, nil

}

func (r *repository) Delete(id int) error {
	err := r.database.Load(&storage)
	if err != nil {
		return err
	}
	for k, v := range storage.Seller {
		if v.ID == id {
			storage.Seller = append(storage.Seller[:k], storage.Seller[k+1:]...)
			err = r.database.Save(&storage)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return customerrors.ErrorInvalidID
}

func (r *repository) GetId(indice int) (models.Seller, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return models.Seller{}, err
	}
	for _, v := range storage.Seller {
		if v.ID == indice {
			return v, nil
		}
	}
	return models.Seller{}, customerrors.ErrorInvalidID
}

func (r *repository) CheckCid(cid int) (models.Seller, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return models.Seller{}, err
	}
	for _, v := range storage.Seller {
		if v.Cid == cid {
			return models.Seller{}, customerrors.ErrorConflict
		}
	}
	return models.Seller{}, nil
}

func (r *repository) Update(newValues Seller, id int) (models.Seller, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return models.Seller{}, err
	}
	for k, v := range storage.Seller {
		if v.ID == id {
			if newValues.Address != "" {
				v.Address = newValues.Address
			}
			if newValues.Cid != 0 {
				v.Cid = newValues.Cid
			}
			if newValues.Company_name != "" {
				v.Company_name = newValues.Company_name
			}
			if newValues.Telephone != "" {
				v.Telephone = newValues.Telephone
			}
			storage.Seller[k] = v
			err = r.database.Save(&storage)
			if err != nil {
				return models.Seller{}, err
			}
			return v, nil
		}
	}
	return models.Seller{}, customerrors.ErrorInvalidID
}

func (r *repository) LastID() (int, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return 0, err
	}
	var lastID = 0
	lastID = storage.LastID
	return lastID, nil
}

func (r *repository) Store(sel models.Seller) (models.Seller, error) {
	err := r.database.Load(&storage)
	if err != nil {
		return models.Seller{}, err
	}
	_, err = r.CheckCid(sel.Cid)
	if err != nil {
		return models.Seller{}, err
	}
	p := models.Seller{ID: sel.ID, Cid: sel.Cid, Company_name: sel.Company_name, Address: sel.Address, Telephone: sel.Telephone}
	storage.Seller = append(storage.Seller, p)
	storage.LastID = storage.LastID + 1
	err = r.database.Save(&storage)
	if err != nil {
		return models.Seller{}, err
	}
	return p, nil
}

func Populate() {
	storage = models.Sellers{Seller: []models.Seller{
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
	},
	}
	storage.LastID = 3
}

func Clean() {
	storage = models.Sellers{}
}
