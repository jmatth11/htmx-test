package db

import (
	"errors"
	"htmx-test/models"
)

type Table interface {
    Add(item *models.Item) error
    All() ([]*models.Item, error)
    Get(idx int) (*models.Item, error)
    Update(idx int, item *models.Item) error
}

type internalDB struct {
    data []*models.Item
}

func NewDB() Table {
    return &internalDB {
        data: make([]*models.Item, 0),
    }
}

func (d *internalDB) All() ([]*models.Item, error) {
    return d.data, nil
}

func (d *internalDB) Add(item *models.Item) error {
    item.Id = len(d.data)
    d.data = append(d.data, item)
    return nil
}

func (d *internalDB) Get(idx int) (*models.Item, error) {
    if idx > len(d.data) || idx < 0 {
        return nil, errors.New("item does not exist")
    }
    return d.data[idx], nil
}

func (d *internalDB) Update(idx int, item *models.Item) error {
    if idx > len(d.data) || idx < 0 {
        return errors.New("item does not exist")
    }
    d.data[idx] = item
    return nil
}
