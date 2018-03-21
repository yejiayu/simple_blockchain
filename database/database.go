package database

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Interface interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
}

func New(path string) (Interface, error) {
	var err error
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &database{db}, nil
}

type database struct {
	db *leveldb.DB
}

func (d *database) Get(key []byte) ([]byte, error) {
	return d.db.Get(key, nil)
}

func (d *database) Put(key []byte, value []byte) error {
	return d.db.Put(key, value, nil)
}
