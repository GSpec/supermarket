package supermarket

// Discounter is able to provide a discount based on a items in a Checkout.
type Discounter interface {
	Discount(c Checkout) int
}
