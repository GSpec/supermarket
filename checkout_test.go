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
