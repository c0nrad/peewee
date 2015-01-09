package main

import "testing"

func TestIsRoyalFlush(t *testing.T) {
	h1 := []Card{Card{Ten, Hearts}, Card{Jack, Hearts}, Card{Queen, Hearts}, Card{King, Hearts}, Card{Ace, Hearts}}
	h2 := []Card{Card{Ten, Clubs}, Card{Jack, Hearts}, Card{Queen, Hearts}, Card{King, Hearts}, Card{Ace, Hearts}}

	if !IsRoyalFlush(h1) {
		t.Error("Failed to corretly determine royal flush")
	}

	if IsRoyalFlush(h2) {
		t.Error("Didn't recognize royal flush")
	}
}

func TestIsStraightFlush(t *testing.T) {
	royal := []Card{Card{Ten, Hearts}, Card{Jack, Hearts}, Card{Queen, Hearts}, Card{King, Hearts}, Card{Ace, Hearts}}
	h2 := []Card{Card{Ace, Hearts}, Card{Two, Hearts}, Card{Three, Hearts}, Card{Four, Hearts}, Card{Five, Hearts}}
	h3 := []Card{Card{Ace, Diamonds}, Card{Two, Hearts}, Card{Three, Hearts}, Card{Four, Hearts}, Card{Five, Hearts}}
	h4 := []Card{Card{Seven, Hearts}, Card{Two, Hearts}, Card{Three, Hearts}, Card{Four, Hearts}, Card{Five, Hearts}}

	if !IsStraightFlush(royal) {
		t.Error("Failed to reocognize royal flush as straight flush")
	}

	if !IsStraightFlush(h2) {
		t.Error("Failed to recognize straight flush starting with ACE as straight flush")
	}

	if IsStraightFlush(h3) {
		t.Error("Incorrectly identified straight flush")
	}

	if IsStraightFlush(h4) {
		t.Error("Incorrectly identified straight flush")
	}

}

func TestIsFourOfAKind(t *testing.T) {
	h1 := []Card{Card{Two, Hearts}, Card{Two, Clubs}, Card{Two, Spades}, Card{Two, Diamonds}}
	h2 := []Card{Card{Three, Hearts}, Card{Two, Clubs}, Card{Two, Spades}, Card{Two, Diamonds}}

	if !IsFourOfAKind(h1) {
		t.Error("Failed to recognize a four of a kind")
	}

	if IsFourOfAKind(h2) {
		t.Error("Incorrectly identified a four of a kind")
	}
}

func TestIsFullHouse(t *testing.T) {
	h1 := []Card{Card{Two, Hearts}, Card{Two, Diamonds}, Card{Two, Clubs}, Card{Three, Diamonds}, Card{Three, Spades}}
	if !IsFullHouse(h1) {
		t.Error("Failed to recognize a full house")
	}

	h2 := []Card{Card{Two, Hearts}, Card{Two, Clubs}, Card{Two, Spades}, Card{Ace, Diamonds}}
	if IsFullHouse(h2) {
		t.Error("Incorrectly identified a Full House")
	}
}

func TestIsFlush(t *testing.T) {
	h1 := []Card{Card{Two, Hearts}, Card{Three, Hearts}, Card{Ace, Hearts}, Card{King, Hearts}, Card{Nine, Hearts}}

	if !IsFlush(h1) {
		t.Error("Failed to recognize flush")
	}

	h2 := []Card{Card{Two, Diamonds}, Card{Three, Hearts}, Card{Ace, Hearts}, Card{King, Hearts}, Card{Nine, Hearts}}
	if IsFlush(h2) {
		t.Error("Incorrectly identified a flush")
	}
}

func TestIsStraight(t *testing.T) {
	royal := []Card{Card{Ten, Hearts}, Card{Jack, Hearts}, Card{Queen, Hearts}, Card{King, Hearts}, Card{Ace, Hearts}}
	if !IsStraight(royal) {
		t.Error("Failed to recognize a straight in a royal")
	}

	h1 := []Card{Card{Ace, Hearts}, Card{Two, Hearts}, Card{Three, Hearts}, Card{Four, Hearts}, Card{Five, Hearts}}
	if !IsStraight(h1) {
		t.Error("Failed to recognize a straight")
	}

	h2 := []Card{Card{Two, Diamonds}, Card{Two, Hearts}, Card{Three, Hearts}, Card{Four, Hearts}, Card{Five, Hearts}}
	if IsStraight(h2) {
		t.Error("Incorrectly identified a straight")
	}
}

func TestThreeOfAKind(t *testing.T) {
	h1 := []Card{Card{Two, Hearts}, Card{Two, Diamonds}, Card{Two, Spades}, Card{King, Diamonds}, Card{Nine, Hearts}}
	if !IsThreeOfAKind(h1) {
		t.Error("Failed to recognize IsThreeOfAKind")
	}

	h2 := []Card{Card{Two, Diamonds}, Card{Two, Hearts}, Card{Ace, Hearts}, Card{Ace, Diamonds}, Card{Nine, Hearts}}
	if IsThreeOfAKind(h2) {
		t.Error("Incorrectly identified a IsThreeOfAKind")
	}
}

func TestTwoPair(t *testing.T) {
	h1 := []Card{Card{Two, Diamonds}, Card{Two, Hearts}, Card{Ace, Hearts}, Card{Ace, Diamonds}, Card{Nine, Hearts}}
	if !IsTwoPair(h1) {
		t.Error("Failed to recognize IsTwoPair")
	}

	h2 := []Card{Card{Two, Diamonds}, Card{Two, Hearts}, Card{Ace, Hearts}, Card{King, Diamonds}, Card{Nine, Hearts}}
	if IsTwoPair(h2) {
		t.Error("Incorrectly identified a IsTwoPair")
	}
}

func TestPair(t *testing.T) {
	h1 := []Card{Card{Two, Diamonds}, Card{Three, Hearts}, Card{Ace, Hearts}, Card{Ace, Diamonds}, Card{Nine, Hearts}}
	if !IsPair(h1) {
		t.Error("Failed to recognize IsPair")
	}

	h2 := []Card{Card{Two, Diamonds}, Card{Three, Hearts}, Card{Four, Hearts}, Card{King, Diamonds}, Card{Nine, Hearts}}
	if IsPair(h2) {
		t.Error("Incorrectly identified a IsPair")
	}
}
