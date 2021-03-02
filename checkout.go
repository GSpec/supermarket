package supermarket

// Checkout provides functionality for scanning SKUs and calculating the total price.
type Checkout struct {
	store *Store
	items []Item
}

func NewCheckout(s *Store) *Checkout {
	return &Checkout{store: s}
}

// Scan adds an item to the checkout.
func (c *Checkout) Scan(sku rune) error {
	item, err := c.store.ChooseItem(sku)

	if err != nil {
		return err
	}

	c.items = append(c.items, *item)
	return nil
}

// GetTotalPrice calculates the price of the items currently in the checkout.
func (c Checkout) GetTotalPrice() int {
	t := 0
	for _, item := range c.items {
		t += item.UnitPrice
	}

	return t
}
