package service

import (
	"fmt"
	c "github.com/gost1k337/wb_demo/subscriber/internal/cache"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"github.com/gost1k337/wb_demo/subscriber/internal/repository"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
)

type OrderService struct {
	cache  *c.Cache
	repo   repository.Order
	logger logging.Logger
}

func NewOrderService(cache *c.Cache, repo repository.Order, logger logging.Logger) *OrderService {
	return &OrderService{
		cache:  cache,
		repo:   repo,
		logger: logger,
	}
}

func (o *OrderService) GetById(id int64) (*entity.Order, error) {
	order, ok := o.cache.Get(id)
	if !ok {
		return nil, fmt.Errorf("order not found")
	}

	return &order, nil
}

func (o *OrderService) Save(order *entity.Order) error {
	id, err := o.repo.Save(order)
	if err != nil {
		return fmt.Errorf("save: %v", err)
	}

	order.Id = id

	o.cache.Set(id, *order)

	return nil
}

func (o *OrderService) LoadCacheFromDb() error {
	orders, err := o.repo.GetAll()
	if err != nil {
		return fmt.Errorf("get all: %w", err)
	}

	for _, order := range orders {
		o.cache.Set(order.Id, *order)
	}

	o.logger.Info("Cache loaded from db...")

	return nil
}
