package subscriber

import (
	"fmt"
	"github.com/gost1k337/wb_demo/subscriber/internal/service"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/nats-io/stan.go"
)

type StanSub struct {
	ss     service.Order
	h      *Handler
	sc     stan.Conn
	logger logging.Logger
}

func NewStanSub(services *service.Services, logger logging.Logger, clusterId, clientId, url string) (*StanSub, error) {
	sc, err := stan.Connect(clusterId, clientId, stan.NatsURL(url))
	if err != nil {
		return nil, fmt.Errorf("conn: %w", err)
	}

	h := NewSubHandler(services.Order, logger)

	sub := &StanSub{
		ss:     services.Order,
		sc:     sc,
		h:      h,
		logger: logger,
	}

	return sub, nil
}

func (ss *StanSub) Close() error {
	err := ss.sc.Close()
	if err != nil {
		return fmt.Errorf("close conn: %w", err)
	}

	return nil
}

func (ss *StanSub) SubscribeToChannel(ch string, opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	sub, err := ss.sc.Subscribe(ch, ss.h.ReceiveOrderHandler, opts...)
	if err != nil {
		return nil, fmt.Errorf("sub: %w", err)
	}

	err = ss.ss.LoadCacheFromDb()
	if err != nil {
		return nil, fmt.Errorf("load cache: %w", err)
	}

	return sub, err
}
