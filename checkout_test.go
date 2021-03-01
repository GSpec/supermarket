package supermarket

import (
	"testing"
)

func TestCheckout_Scan(t *testing.T) {
	type args struct {
		sku rune
	}
	tests := map[string]struct {
		c       *Checkout
		args    args
		wantErr bool
	}{
		"Scan A": {new(Checkout), args{'A'}, false},
		"Scan B": {new(Checkout), args{'B'}, false},
		"Scan C": {new(Checkout), args{'C'}, true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.c.Scan(tt.args.sku)

			if (err != nil) != tt.wantErr {
				t.Errorf("Unwanted error returned: %v", err)
			}
		})
	}
}

func TestCheckout_GetTotalPrice(t *testing.T) {
	tests := map[string]struct {
		c    *Checkout
		skus []rune
		want int
	}{
		"Checkout A":     {new(Checkout), []rune{'A'}, 50},
		"Checkout A B":   {new(Checkout), []rune{'A', 'B'}, 80},
		"Checkout A B C": {new(Checkout), []rune{'A', 'B', 'C'}, 80},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			for _, sku := range tt.skus {
				tt.c.Scan(sku)
			}

			if got := tt.c.GetTotalPrice(); got != tt.want {
				t.Errorf("Incorrect price, wanted %v, got: %v", tt.want, got)
			}
		})
	}
}
