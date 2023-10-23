package service

import (
	"github.com/gost1k337/wb_demo/subscriber/internal/cache"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"github.com/gost1k337/wb_demo/subscriber/internal/repository"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
)

type Order interface {
	GetById(id int64) (*entity.Order, error)
	Save(*entity.Order) error
	LoadCacheFromDb() error
}

type Services struct {
	Order
}

type Deps struct {
	Cache *cache.Cache
	Repos *repository.Repositories
}

func NewServices(deps *Deps, logger logging.Logger) *Services {
	return &Services{
		Order: NewOrderService(deps.Cache, deps.Repos.Order, logger),
	}
}
