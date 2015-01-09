package main

import "testing"

func TestRankToChar(t *testing.T) {
	if RankToChar(Ace) != 'A' {
		t.Error("Failed to convert ace to proper string")
	}
}

func TestSuitToChar(t *testing.T) {
	if SuitToChar(Clubs) != 'C' {
		t.Error("Failed to convert clubs to proper string")
	}
}

// func TestCardString(t *testing.T) {
// 	c := Card{Three, Clubs}
// 	if c.String() != "3C" {
// 		t.Error("Failed to generate correct string from card")
// 	}
// }

func TestCardEquality(t *testing.T) {
	c1 := Card{Ace, Hearts}
	c2 := Card{King, Hearts}

	if c1 == c2 {
		t.Error("Incorrectly said two cards were equal", c1, c2)
	}

	c2.Rank = Ace
	if c1 != c2 {
		t.Error("Incorrectly said two cards were not equal", c1, c2)
	}
}
