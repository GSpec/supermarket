package supermarket

import (
	"fmt"
	"testing"
)

func TestStore_LoadStock(t *testing.T) {
	tests := map[string]struct {
		stock   map[rune]Item
		wantErr string
	}{
		"Valid Stock A":               {map[rune]Item{'A': {1}}, ""},
		"Valid Stock Z":               {map[rune]Item{'Z': {2147483647}}, ""},
		"Invalid Stock SKU Lowercase": {map[rune]Item{'a': {1}}, fmt.Sprintf(errorInvalidSku, 'a')},
		"Invalid Stock SKU Symbol":    {map[rune]Item{'%': {1}}, fmt.Sprintf(errorInvalidSku, '%')},
		"Invalid Stock Price 0":       {map[rune]Item{'Z': {0}}, fmt.Sprintf(errorInvalidPrice, 0)},
		"Invalid Stock Price -1":      {map[rune]Item{'Z': {-1}}, fmt.Sprintf(errorInvalidPrice, -1)},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := new(Store)
			err := s.LoadStock(tt.stock)

			assertErrorCorrect(t, tt.wantErr, err)
		})
	}
}

func TestStore_LoadOffers(t *testing.T) {
	tests := map[string]struct {
		offers  []Discounter
		wantErr string
	}{
		"No Offers":      {[]Discounter{}, ""},
		"Multibuy Offer": {[]Discounter{Multibuy{'A', 2, 10}}, ""},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := new(Store)
			err := s.LoadOffers(tt.offers)

			assertErrorCorrect(t, tt.wantErr, err)
		})
	}
}

func assertErrorCorrect(t *testing.T, wantErr string, err error) {
	switch {
	case err != nil && wantErr == "":
		t.Errorf("Unwanted error returned: %v", err)
	case err != nil && wantErr != err.Error():
		t.Errorf("Wanted error: %v, got: %v", wantErr, err)
	case err == nil && wantErr != "":
		t.Errorf("Wanted error: %v, got: %v", wantErr, err)
	}
}
