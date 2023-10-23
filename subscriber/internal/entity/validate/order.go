package validate

import (
	"fmt"
	"github.com/gost1k337/wb_demo/subscriber/internal/entity"
)

func CheckDelivery(d *entity.Delivery) error {
	if len(d.Name) < 1 {
		return InvalidNameErr
	}

	if d.Phone == "" {
		return fmt.Errorf("no field phone")
	}

	return nil
}

func CheckPayment(p *entity.Payment) error {
	if p.Amount <= 0 {
		return InvalidAmountErr
	}

	if p.Transaction == "" {
		return fmt.Errorf("no field transaction")
	}

	return nil
}

func CheckOrder(order *entity.Order) error {
	if len(order.OrderUID) == 0 {
		return InvalidIDErr
	}

	if err := CheckDelivery(&order.Delivery); err != nil {
		return fmt.Errorf("delivery: %w", err)
	}

	if err := CheckPayment(&order.Payment); err != nil {
		return fmt.Errorf("payment: %w", err)
	}

	for _, i := range order.Items {
		if i.Price < 0 || i.TotalPrice < 0 {
			return InvalidItemPriceErr
		}
	}

	return nil
}
