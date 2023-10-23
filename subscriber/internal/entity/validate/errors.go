package validate

import "errors"

var (
	InvalidIDErr     = errors.New("invalid id")
	InvalidItemPrice = errors.New("invalid item price")
)
