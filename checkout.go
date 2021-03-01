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
		c.items = append(c.items, Item{sku, 50})
		return nil
	default:
		return errors.New(fmt.Sprintf("Could not find SKU: %v", sku))
	}
}
