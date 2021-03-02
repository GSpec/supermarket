package supermarket

import (
	"errors"
	"fmt"
)

const errorSkuNotFound = "Could not find SKU: %c"

type Store struct {
	stock map[rune]Item
}

func (s *Store) LoadStock(stock map[rune]Item) {
	s.stock = stock
}

func (s Store) ChooseItem(sku rune) (*Item, error) {
	if item, ok := s.stock[sku]; ok {
		return &item, nil
	}

	return nil, errors.New(fmt.Sprintf("Could not find SKU: %c", sku))
}
