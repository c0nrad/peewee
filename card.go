package main

import (
	"sort"
	"strconv"
)

const (
	NumberOfRanks     = 13
	NumberOfRanksSafe = 15

	LowAce = 1 //only used for scoring
	Two    = 2
	Three  = 3
	Four   = 4
	Five   = 5
	Six    = 6
	Seven  = 7
	Eight  = 8
	Nine   = 9
	Ten    = 10
	Jack   = 11
	Queen  = 12
	King   = 13
	Ace    = 14
)

const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

var Suits = []int{Clubs, Diamonds, Hearts, Spades}

type Card struct {
	Rank, Suit int
}

type SortableCards []Card

func (h SortableCards) Len() int           { return len(h) }
func (h SortableCards) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h SortableCards) Less(i, j int) bool { return h[i].Rank < h[j].Rank }

func RankToChar(rank int) byte {
	if rank == LowAce {
		return 'A'
	}

	if rank >= 2 && rank <= 9 {
		return byte(rank + '0')
	}

	if rank == Ten {
		return 'T'
	} else if rank == Jack {
		return 'J'
	} else if rank == Queen {
		return 'Q'
	} else if rank == King {
		return 'K'
	} else if rank == Ace {
		return 'A'
	}

	return '?'
}

func RankToString(rank int) string {
	if rank == LowAce {
		return "Ace"
	}

	if rank >= 2 && rank <= 9 {
		return strconv.Itoa(rank)
	}

	if rank == Ten {
		return "10"
	} else if rank == Jack {
		return "Jack"
	} else if rank == Queen {
		return "Queen"
	} else if rank == King {
		return "King"
	} else if rank == Ace {
		return "Ace"
	}

	return "?" + strconv.Itoa(rank) + "?"
}

func SuitToChar(suit int) byte {
	switch suit {
	case Clubs:
		return 'C'
	case Diamonds:
		return 'D'
	case Hearts:
		return 'H'
	case Spades:
		return 'S'
	default:
		return '?'
	}
}

func SuitToString(suit int) string {
	switch suit {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return "?" + strconv.Itoa(suit) + "?"
	}
}

func Contains(cards []Card, c Card) bool {
	for _, t := range cards {
		if t == c {
			return true
		}
	}
	return false
}

func ContainsRank(cards []Card, rank int) bool {
	for _, c := range cards {
		if c.Rank == rank {
			return true
		}
	}
	return false
}

func ContainsAll(cards, needs []Card) bool {
	for _, n := range needs {
		if !Contains(cards, n) {
			return false
		}
	}
	return true
}

func GetRanks(cards []Card, rank int) (out []Card) {
	for _, c := range cards {
		if c.Rank == rank {
			out = append(out, c)
		}
	}
	return out
}

func GetSuits(cards []Card, suit int) (out []Card) {
	for _, c := range cards {
		if c.Suit == suit {
			out = append(out, c)
		}
	}
	return out
}

// Little first
func SortByRank(cards []Card) []Card {
	sort.Sort(SortableCards(cards))
	return cards
}

func (c Card) String() string {
	// out := []byte{}
	// out = append(out, RankToChar(c.Rank))
	// out = append(out, SuitToChar(c.Suit))
	// return string(out)

	return "(" + RankToString(c.Rank) + " " + SuitToString(c.Suit) + ")"
}
