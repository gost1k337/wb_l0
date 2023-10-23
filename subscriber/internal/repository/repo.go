package repository

import (
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"github.com/gost1k337/wb_demo/subscriber/internal/repository/psql"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/gost1k337/wb_demo/subscriber/pkg/postgres"
)

type Order interface {
	GetAll() ([]*entity.Order, error)
	Save(*entity.Order) (int64, error)
}

type Repositories struct {
	Order
}

func NewRepositories(pg *postgres.Postgres, logger logging.Logger) *Repositories {
	return &Repositories{
		Order: psql.NewOrderRepo(pg, logger),
	}
}
