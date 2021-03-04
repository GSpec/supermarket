package supermarket

import (
	"fmt"
	"testing"
)

func TestMultibuy_NewMultibuy(t *testing.T) {
	tests := map[string]struct {
		sku      rune
		quantity int
		discount int
		wantErr  string
	}{
		"Valid Offer SKU A":   {sku: 'A', quantity: 3, discount: 20, wantErr: ""},
		"Valid Offer SKU B":   {sku: 'B', quantity: 2, discount: 12, wantErr: ""},
		"Invalid Quantity 1":  {sku: 'A', quantity: 1, discount: 10, wantErr: fmt.Sprintf(errorInvalidMultibuyQuantity, 1)},
		"Invalid Quantity 0":  {sku: 'A', quantity: 0, discount: 10, wantErr: fmt.Sprintf(errorInvalidMultibuyQuantity, 0)},
		"Invalid Quantity -1": {sku: 'A', quantity: -1, discount: 10, wantErr: fmt.Sprintf(errorInvalidMultibuyQuantity, -1)},
		"Invalid Discount 0":  {sku: 'A', quantity: 3, discount: 0, wantErr: fmt.Sprintf(errorInvalidMultibuyDiscount, 0)},
		"Invalid Discount -1": {sku: 'A', quantity: 4, discount: -1, wantErr: fmt.Sprintf(errorInvalidMultibuyDiscount, -1)},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			offer, err := NewMultibuy(tt.sku, tt.quantity, tt.discount)

			switch {
			case err != nil && tt.wantErr == "":
				t.Errorf("Unwanted error returned: %v", err)
			case err != nil && tt.wantErr != err.Error():
				t.Errorf("Wanted error: %v, got: %v", tt.wantErr, err)
			case err == nil && tt.wantErr != "":
				t.Errorf("Wanted error: %v, got: %v", tt.wantErr, err)
			case err == nil && offer.quantity != tt.quantity:
				t.Errorf("Quantity not set correctly, wanted: %v, got: %v", tt.quantity, offer.quantity)
			case err == nil && offer.discountAmount != tt.discount:
				t.Errorf("Quantity not set correctly, wanted: %v, got: %v", tt.discount, offer.discountAmount)
			}
		})
	}
}

func TestMultibuy_calculateDiscount(t *testing.T) {
	tests := map[string]struct {
		offer            Multibuy
		checkoutSku      rune
		checkoutQuantity int
		want             int
	}{
		"Buy 3 Discount 20 Once":  {offer: Multibuy{'A', 3, 20}, checkoutSku: 'A', checkoutQuantity: 3, want: 20},
		"Buy 2 Discount 15 Once":  {offer: Multibuy{'B', 2, 15}, checkoutSku: 'B', checkoutQuantity: 2, want: 15},
		"Zero Checkout Quantity":  {offer: Multibuy{'A', 3, 15}, checkoutSku: 'A', checkoutQuantity: 0, want: 0},
		"Buy 4 Discount 30 Twice": {offer: Multibuy{'A', 2, 30}, checkoutSku: 'A', checkoutQuantity: 4, want: 60},
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
