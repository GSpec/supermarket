package supermarket

// Checkout provides functionality for scanning SKUs and calculating the total price.
type Checkout struct {
	store     *Store
	lineItems []LineItem
}

// LineItem is an item in a Checkout.
type LineItem struct {
	sku  rune
	item Item
}

// NewCheckout creates an instance of a Checkout using the Store provided.
func NewCheckout(s *Store) *Checkout {
	return &Checkout{store: s}
}

// Scan adds an item to the checkout, or returns an error if not found.
func (c *Checkout) Scan(sku rune) error {
	item, err := c.store.ChooseItem(sku)

	if err != nil {
		return err
	}

	c.lineItems = append(c.lineItems, LineItem{sku, *item})
	return nil
}

// GetTotalPrice calculates the price of the items currently in the checkout.
func (c Checkout) GetTotalPrice() int {
	t := 0
	for _, li := range c.lineItems {
		t += li.item.UnitPrice
	}

	for _, discounter := range c.store.offers {
		t -= discounter.Discount(c)
	}

	return t
}

func (c Checkout) getCountOf(sku rune) int {
	t := 0
	for _, li := range c.lineItems {
		if li.sku == sku {
			t++
		}
	}

	return t
}
