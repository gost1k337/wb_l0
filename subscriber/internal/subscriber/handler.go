package subscriber

import (
	"encoding/json"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity/validate"
	"github.com/gost1k337/wb_demo/subscriber/internal/service"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/nats-io/stan.go"
)

type Handler struct {
	ss     service.Order
	logger logging.Logger
}

func NewSubHandler(ss service.Order, logger logging.Logger) *Handler {
	return &Handler{
		ss:     ss,
		logger: logger,
	}
}

func (h *Handler) ReceiveOrderHandler(msg *stan.Msg) {
	h.logger.Info("Received new message on channel")

	b := msg.Data
	order := &entity.Order{}
	if err := json.Unmarshal(b, order); err != nil {
		h.logger.Errorf("read data: %v", err)
		return
	}

	if err := validate.ValidateOrder(order); err != nil {
		h.logger.Errorf("validate: %v", err)
		return
	}

	err := h.ss.Save(order)
	if err != nil {
		h.logger.Errorf("order: %v", err)
		return
	}
}
