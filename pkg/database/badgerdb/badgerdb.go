package badgerdb

import (
	badger "github.com/dgraph-io/badger/v4"
	"github.com/lefes/discord-message-scheduler/internal/config"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

func NewClient(cfg config.DBConfig) (*badger.DB, error) {
	options := badger.DefaultOptions(cfg.Path).WithLogger(logger.Log)
	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	return db, nil
}
