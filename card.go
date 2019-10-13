package main

import (
"errors"
"fmt"
)

type Card struct {
	rank rune
	suit rune
}


func cardAt(num int) (*Card, error) {
	if num < 0 || 51 < num {
		return nil, errors.New("Index out if range")
	}

	rank := []rune{'2', '3', '4', '5', '6', '7', '8', '9', '0', 'J', 'Q', 'K', 'A'}
	suit := []rune{'C', 'D', 'H', 'S'}

	indexR := num % 13
	indexS := num / 13

	return &Card{rank[indexR], suit[indexS]}, nil
}


func (card Card) String() string {
	return fmt.Sprintf("%c%c", card.rank, card.suit)
}