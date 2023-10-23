package validate

import "github.com/gost1k337/wb_demo/subscriber/internal/entity"

func ValidateOrder(order *entity.Order) error {
	if len(order.OrderUID) == 0 {
		return InvalidIDErr
	}

	for _, i := range order.Items {
		if i.Price < 0 || i.TotalPrice < 0 {
			return InvalidItemPrice
		}
	}

	return nil
}
