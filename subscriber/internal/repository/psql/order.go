package psql

import (
	"encoding/json"
	"fmt"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/gost1k337/wb_demo/subscriber/pkg/postgres"
)

type OrderRepo struct {
	db     *postgres.Postgres
	logger logging.Logger
}

func NewOrderRepo(db *postgres.Postgres, logger logging.Logger) *OrderRepo {
	return &OrderRepo{
		db:     db,
		logger: logger,
	}
}

func (o *OrderRepo) GetAll() ([]*entity.Order, error) {
	query := `SELECT * from orders`

	rows, err := o.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	orders := make([]*entity.Order, 0)

	for rows.Next() {
		deliveryJson := json.RawMessage{}
		paymentJson := json.RawMessage{}
		itemsJson := json.RawMessage{}
		order := &entity.Order{}

		if err := rows.Scan(
			&order.Id,
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&deliveryJson,
			&paymentJson,
			&itemsJson,
			&order.Locale,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SMID,
			&order.DateCreated,
			&order.OofShard,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		delivery := entity.Delivery{}
		payment := entity.Payment{}
		items := make([]entity.Item, 0)

		if err := json.Unmarshal(deliveryJson, &delivery); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}

		if err := json.Unmarshal(paymentJson, &payment); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}

		if err := json.Unmarshal(itemsJson, &items); err != nil {
			return nil, fmt.Errorf("unmarshal: %w", err)
		}

		order.Delivery = delivery
		order.Payment = payment
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (o *OrderRepo) Save(order *entity.Order) (int64, error) {
	payment, err := json.Marshal(order.Payment)
	if err != nil {
		return 0, fmt.Errorf("marshal: %w", err)
	}

	delivery, err := json.Marshal(order.Delivery)
	if err != nil {
		return 0, fmt.Errorf("marshal: %w", err)
	}

	items, err := json.Marshal(order.Items)
	if err != nil {
		return 0, fmt.Errorf("marshal: %w", err)
	}

	query := `INSERT INTO orders (order_uid, track_number, entry, deliveries, payments, items, locale, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`

	var id int64
	_ = o.db.QueryRow(query, order.OrderUID, order.TrackNumber,
		order.Entry, delivery, payment, items, order.Locale,
		order.CustomerID, order.DeliveryService, order.ShardKey,
		order.SMID, order.DateCreated, order.OofShard,
	).Scan(&id)

	return id, nil
}
