package supermarket

// Multibuy offer to get a discount when multiple items are bought.
type Multibuy struct {
	sku            rune
	quantity       int
	discountAmount int
}

func (o Multibuy) Discount(c Checkout) int {
	skuCount := c.getCountOf(o.sku)
	return o.calculateDiscount('A', skuCount)
}

func (o Multibuy) calculateDiscount(sku rune, quantity int) int {
	if o.sku != sku || o.quantity > quantity {
		return 0
	}

	return o.discountAmount
}
