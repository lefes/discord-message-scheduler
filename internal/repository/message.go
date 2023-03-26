package repository

import (
	"github.com/dgraph-io/badger/v4"
)

type MessageRepo struct {
	db *badger.DB
}

func NewMessageRepository(db *badger.DB) *MessageRepo {
	return &MessageRepo{
		db: db,
	}
}

func (r *MessageRepo) SetKey(key string, value string) error {
	err := r.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepo) GetKey(key string) (string, error) {
	var valCopy []byte
	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return "", err
	}
	return string(valCopy), nil
}

func (r *MessageRepo) DeleteKey(key string) error {
	err := r.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
