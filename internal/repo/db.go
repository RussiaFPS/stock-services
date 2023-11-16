package repo

import (
	"stock-services/pkg/postgres"
)

type Storage interface {
}

type DbRepo struct {
	db postgres.Db
}

func NewRepository(db postgres.Db) Storage {
	return &DbRepo{db}
}
