package validate

import "errors"

var (
	InvalidIDErr        = errors.New("invalid id")
	InvalidItemPriceErr = errors.New("invalid item price")
	InvalidNameErr      = errors.New("invalid name")
	InvalidAmountErr    = errors.New("invalid amount")
)
