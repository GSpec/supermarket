package supermarket

import (
	"errors"
	"fmt"
	"unicode"
)

const errorInvalidSku = "Could not load item with SKU: %c"
const errorInvalidPrice = "Could not load item with UnitPrice: %v"
const errorSkuNotFound = "Could not find SKU: %c"

type Store struct {
	stock  map[rune]Item
	offers []Discounter
}

func (s *Store) LoadStock(stock map[rune]Item) error {
	for sku, item := range stock {
		switch {
		case !unicode.IsUpper(sku):
			return errors.New(fmt.Sprintf(errorInvalidSku, sku))
		case item.UnitPrice < 1:
			return errors.New(fmt.Sprintf(errorInvalidPrice, item.UnitPrice))
		}
	}
	s.stock = stock
	return nil
}

func (s *Store) LoadOffers(offers []Discounter) error {
	s.offers = offers
	return nil
}

func (s Store) ChooseItem(sku rune) (*Item, error) {
	if item, ok := s.stock[sku]; ok {
		return &item, nil
	}

	return nil, errors.New(fmt.Sprintf(errorSkuNotFound, sku))
}
