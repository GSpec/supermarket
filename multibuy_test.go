package supermarket

import (
	"testing"
)

func TestMultibuy_calculateDiscount(t *testing.T) {
	tests := map[string]struct {
		offer            Multibuy
		checkoutSku      rune
		checkoutQuantity int
		want             int
	}{
		"Buy 2 Discount 10 Once": {offer: Multibuy{'A', 3, 20}, checkoutSku: 'A', checkoutQuantity: 3, want: 20},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			discount := tt.offer.calculateDiscount(tt.checkoutSku, tt.checkoutQuantity)

			if discount != tt.want {
				t.Errorf("Wanted discount amount: %v, got: %v", tt.want, discount)
			}
		})
	}
}
