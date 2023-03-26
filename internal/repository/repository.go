package repository

import "github.com/dgraph-io/badger/v4"

type Repositories struct {
	Message Message
}

func NewRepositories(db *badger.DB) *Repositories {
	return &Repositories{
		Message: NewMessageRepository(db),
	}
}

type Message interface {
	SetKey(key string, value string) error
	GetKey(key string) (string, error)
	DeleteKey(key string) error
}
