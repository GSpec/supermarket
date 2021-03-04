package supermarket

import (
	"errors"
	"fmt"
)

const errorInvalidMultibuyQuantity = "Invalid Multibuy Quantity: %v"
const errorInvalidMultibuyDiscount = "Invalid Multibuy Discount: %v"

// Multibuy offer to get a discount when multiple items are bought.
type Multibuy struct {
	sku            rune
	quantity       int
	discountAmount int
}

func NewMultibuy(sku rune, quantity int, discountAmount int) (*Multibuy, error) {
	switch {
	case quantity < 2:
		return nil, errors.New(fmt.Sprintf(errorInvalidMultibuyQuantity, quantity))
	case discountAmount <= 0:
		return nil, errors.New(fmt.Sprintf(errorInvalidMultibuyDiscount, discountAmount))
	}
	return &Multibuy{sku, quantity, discountAmount}, nil
}

func (o Multibuy) Discount(c Checkout) int {
	skuCount := c.getCountOf(o.sku)
	return o.calculateDiscount(o.sku, skuCount)
}

func (o Multibuy) calculateDiscount(sku rune, quantity int) int {
	if o.sku != sku || o.quantity > quantity {
		return 0
	}

	offerApplyCount := quantity / o.quantity
	return o.discountAmount * offerApplyCount
}
