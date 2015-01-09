package main

import "testing"

func TestDeckDraw(t *testing.T) {
	d := NewDeck()
	d, c := Draw(d)

	if len(d) != 51 {
		t.Error("Deck should be 51 cards after drawing one")
	}

	for _, c1 := range d {
		if c1 == c {
			t.Error("Card should not be in deck after drawing")
		}
	}
}
