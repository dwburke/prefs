package storage

import (
	"github.com/spf13/viper"

	"github.com/dwburke/prefs/storage/common"
	"github.com/dwburke/prefs/storage/memory"
	"github.com/dwburke/prefs/storage/meta"
	"github.com/dwburke/prefs/storage/mysql"
)

type Storage struct {
	base meta.Storage
}

func New() (*Storage, error) {

	storage_type := viper.GetString("prefs.storage.type")

	var base meta.Storage
	var err error

	switch storage_type {
	case "memory":
		base, err = memory.New()
	case "mysql":
		base, err = mysql.New()
	default:
		err = common.ErrInvalidDatabase
	}

	if err != nil {
		return nil, err
	}

	return &Storage{base}, nil
}

func (store *Storage) Set(key string, value []byte) error {
	return store.base.Set(key, value)
}

func (store *Storage) Get(key string) ([]byte, error) {
	return store.base.Get(key)
}

func (store *Storage) Delete(key string) error {
	return store.base.Delete(key)
}

func (store *Storage) Close() error {
	return store.base.Close()
}
