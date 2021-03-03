package supermarket

import (
	"fmt"
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
	testStore.LoadOffers([]Discounter{
		Multibuy{sku: 'A', quantity: 3, discountAmount: 20},
	})
	os.Exit(m.Run())
}

func TestCheckout_Scan(t *testing.T) {
	tests := map[string]struct {
		sku     rune
		wantErr string
	}{
		"Scan A": {'A', ""},
		"Scan B": {'B', ""},
		"Scan E": {'E', fmt.Sprintf(errorSkuNotFound, 'E')},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCheckout(testStore)
			err := c.Scan(tt.sku)

			switch {
			case err != nil && tt.wantErr == "":
				t.Errorf("Unwanted error returned: %v", err)
			case err != nil && tt.wantErr != err.Error():
				t.Errorf("Wanted error: %v, got: %v", tt.wantErr, err)
			case err == nil && tt.wantErr != "":
				t.Errorf("Wanted error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}

func TestCheckout_GetTotalPrice(t *testing.T) {
	tests := map[string]struct {
		skus []rune
		want int
	}{
		"Checkout A":     {[]rune{'A'}, 50},
		"Checkout A B":   {[]rune{'A', 'B'}, 80},
		"Checkout A B E": {[]rune{'A', 'B', 'E'}, 80},
		"Checkout None":  {[]rune{}, 0},
		"Checkout A A A": {[]rune{'A', 'A', 'A'}, 130},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCheckout(testStore)
			for _, sku := range tt.skus {
				c.Scan(sku)
			}

			if got := c.GetTotalPrice(); got != tt.want {
				t.Errorf("Incorrect price, wanted %v, got: %v", tt.want, got)
			}
		})
	}
}
