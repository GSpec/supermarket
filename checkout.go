package supermarket

import (
	"errors"
	"fmt"
)

// Checkout provides functionality for scanning SKUs and calculating the total price.
type Checkout struct {
	items []Item
}

// Scan adds an item to the checkout.
func (c *Checkout) Scan(sku rune) error {
	switch sku {
	case 'A':
		c.items = append(c.items, Item{sku, 50})
		return nil
	case 'B':
		c.items = append(c.items, Item{sku, 30})
		return nil
	default:
		return errors.New(fmt.Sprintf("Could not find SKU: %v", sku))
	}
}

// GetTotalPrice calculates the price of the items currently in the checkout.
func (c Checkout) GetTotalPrice() int {
	t := 0
	for _, item := range c.items {
		t += item.UnitPrice
	}

	return t
}
