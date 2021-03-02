package supermarket

import (
	"os"
	"testing"
)

var testStore *Store

func TestMain(m *testing.M) {
	testStore = new(Store)
	testStore.LoadStock(map[rune]Item{
		'A': {50},
		'B': {30},
		'C': {20},
		'D': {15},
	})
	os.Exit(m.Run())
}

func TestCheckout_Scan(t *testing.T) {
	type args struct {
		sku rune
	}
	tests := map[string]struct {
		c       *Checkout
		args    args
		wantErr bool
	}{
		"Scan A": {NewCheckout(testStore), args{'A'}, false},
		"Scan B": {NewCheckout(testStore), args{'B'}, false},
		"Scan E": {NewCheckout(testStore), args{'E'}, true},
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
		"Checkout A":     {NewCheckout(testStore), []rune{'A'}, 50},
		"Checkout A B":   {NewCheckout(testStore), []rune{'A', 'B'}, 80},
		"Checkout A B E": {NewCheckout(testStore), []rune{'A', 'B', 'E'}, 80},
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
