package main

import "math"

const (
	RoyalFlush = iota
	FourOfAKind
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func IsRoyalFlush(cards []Card) bool {
	for _, suit := range Suits {
		royal := []Card{Card{Ten, suit}, Card{Jack, suit}, Card{Queen, suit}, Card{King, suit}, Card{Ace, suit}}

		if ContainsAll(cards, royal) {
			return true
		}
	}
	return false
}

func IsStraightFlush(cards []Card) bool {
	// If there's an ace, I create a low ace too
	aces := GetRanks(cards, Ace)
	for _, a := range aces {
		lowAce := Card{LowAce, a.Suit}
		cards = append(cards, lowAce)
	}

	// check each card if it's the start of a straight flush
	for _, c := range cards {
		for rank := c.Rank; rank <= c.Rank+4; rank++ {
			nextCard := Card{rank, c.Suit}
			if !Contains(cards, nextCard) {
				break
			}
			if rank == c.Rank+4 {
				return true
			}
		}
	}
	return false
}

func IsFourOfAKind(cards []Card) bool {
	for _, c := range cards {
		kind := []Card{Card{c.Rank, Hearts}, Card{c.Rank, Spades}, Card{c.Rank, Diamonds}, Card{c.Rank, Clubs}}
		if ContainsAll(cards, kind) {
			return true
		}
	}
	return false
}

func IsFullHouse(cards []Card) bool {
	hasTwoPair := false
	hasThreePair := false
	for _, c := range cards {
		sameRanks := GetRanks(cards, c.Rank)
		if len(sameRanks) == 2 {
			hasTwoPair = true
		} else if len(sameRanks) == 3 {
			hasThreePair = true
		}
	}

	if hasTwoPair && hasThreePair {
		return true
	}
	return false
}

func IsFlush(cards []Card) bool {
	for _, suit := range Suits {
		if len(GetSuits(cards, suit)) == 5 {
			return true
		}
	}
	return false
}

func IsStraight(cards []Card) bool {
	aces := GetRanks(cards, Ace)
	for _, a := range aces {
		lowAce := Card{LowAce, a.Suit}
		cards = append(cards, lowAce)
	}

	for _, c := range cards {
		for rank := c.Rank; rank <= c.Rank+4; rank++ {
			if !ContainsRank(cards, rank) {
				break
			}
			if rank == c.Rank+4 {
				return true
			}
		}
	}
	return false

}

func IsThreeOfAKind(cards []Card) bool {
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 3 {
			return true
		}
	}
	return false
}

func IsTwoPair(cards []Card) bool {
	prevRank := 0
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 2 {
			if prevRank == 0 {
				prevRank = c.Rank
			} else if prevRank != 0 && prevRank != c.Rank {
				return true
			}
		}
	}
	return false
}

func IsPair(cards []Card) bool {
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 2 {
			return true
		}
	}
	return false
}

func IntPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func RoyalFlushScore(cards []Card) int {
	for _, suit := range Suits {
		royal := []Card{Card{Ten, suit}, Card{Jack, suit}, Card{Queen, suit}, Card{King, suit}, Card{Ace, suit}}

		if ContainsAll(cards, royal) {
			return 1
		}
	}
	return 0
}

//Do multiple players have a straight flush? If yes, the winner is the one
//with the highest card. If multiple people share the highest card (obviously
//in a different suit) they split the pot. (Note: Royal flush is excluded
//because it's just a special straight flush that no one else can beat.)
func StraightFlushScore(cards []Card) int {
	// If there's an ace, I create a low ace too
	aces := GetRanks(cards, Ace)
	for _, a := range aces {
		lowAce := Card{LowAce, a.Suit}
		cards = append(cards, lowAce)
	}

	// check each card if it's the start of a straight flush
	for _, c := range cards {
		for rank := c.Rank; rank <= c.Rank+4; rank++ {
			nextCard := Card{rank, c.Suit}
			if !Contains(cards, nextCard) {
				break
			}
			if rank == c.Rank+4 {
				return rank
			}
		}
	}
	return 0
}

//Do multiple players have 4 of a kind? If yes, the one with the highest
//'set of 4' is the winner. If multiple players have the highest set of 4
//(which is not achievable with a standard poker deck, but is with a double
//deck or community cards), the one with the highest kicker (highest card not
//in the set of 4) is the winner. If this card is the same, they split the pot.
func FourOfAKindScore(cards []Card) int {
	winnerRank := 0
	for _, c := range cards {
		kind := []Card{Card{c.Rank, Hearts}, Card{c.Rank, Spades}, Card{c.Rank, Diamonds}, Card{c.Rank, Clubs}}
		if ContainsAll(cards, kind) {
			winnerRank = c.Rank
		}
	}
	if winnerRank == 0 {
		return 0
	}

	//Calulate kicker
	offCard := 0
	for _, c := range cards {
		if c.Rank != winnerRank {
			offCard = c.Rank
		}
	}
	return NumberOfRanks*winnerRank + offCard
}

//Do multiple players have full houses? If yes, then keeping in mind that a
//full house is a 3-set and a 2-set, the player with the highest 3-set wins
//the pot. If multiple players share the highest 3-set (which isn't possible
//without community cards like in hold 'em, or a double deck) then the player
//with the highest 2-set is the winner. If the 2-set and 3-set is the same,
//those players split the pot.
func FullHouseScore(cards []Card) int {
	twoPairRank := 0
	threePairRank := 0
	for _, c := range cards {
		sameRanks := GetRanks(cards, c.Rank)
		if len(sameRanks) == 2 {
			twoPairRank = c.Rank
		} else if len(sameRanks) == 3 {
			threePairRank = c.Rank
		}
	}

	if twoPairRank != 0 && threePairRank != 0 {
		return threePairRank*NumberOfRanksSafe + twoPairRank
	}
	return 0
}

//Do multiple players have a flush? If yes, the player with a flush with
//the highest unique card is the winner. This hand is similar to 'high card'
//resolution, where each card is effectively a kicker. Note that a flush
//requires the same suit, not just color. While the colors used on the suit
//are red and black, two each, there's nothing to that connection. A club is
//no more similar to a spade than it is to a heart - only suit matters. The
//colors are red and black for historical purposes and so the same deck can
//be played for other games where that might matter.
func FlushScore(cards []Card) int {
	for _, suit := range Suits {
		if len(GetSuits(cards, suit)) == 5 {
			cards = SortByRank(cards)
			score := 0
			for i, c := range cards {
				score += IntPow(NumberOfRanksSafe, i) * c.Rank
			}
			return score
		}
	}
	return 0
}

//Do multiple players have straights? If so, the player with the highest
//straight wins. (a-2-3-4-5 is the lowest straight, while 10-j-q-k-a is the
//highest straight.) If multiple players share the highest straight, they
//split the pot.
func StraightScore(cards []Card) int {
	aces := GetRanks(cards, Ace)
	for _, a := range aces {
		lowAce := Card{LowAce, a.Suit}
		cards = append(cards, lowAce)
	}

	for _, c := range cards {
		for rank := c.Rank; rank <= c.Rank+4; rank++ {
			if !ContainsRank(cards, rank) {
				break
			}
			if rank == c.Rank+4 {
				return rank
			}
		}
	}
	return 0

}

//Do multiple players have 3 of a kind? If yes, the player with the highest
//3-set wins the pot. If multiple players have the highest 3-set, the player
//with the highest kicker wins the pot. If multiple players tie for highest
//3-set and highest kicker, the player with the next highest kicker wins the
//pot. If the players tie for the highest 3-set, highest kicker, and highest
//second kicker, the players split the pot.
func ThreeOfAKindScore(cards []Card) int {
	threePairRank := 0
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 3 {
			threePairRank = c.Rank
			break
		}
	}
	if threePairRank == 0 {
		return 0
	}

	kickerCards := []Card{}
	for _, c := range cards {
		if c.Rank != threePairRank {
			kickerCards = append(kickerCards, c)
		}
	}

	kickerCards = SortByRank(kickerCards)
	return IntPow(NumberOfRanksSafe, 2)*threePairRank + IntPow(NumberOfRanksSafe, 1)*kickerCards[1].Rank + kickerCards[0].Rank
}

//Do multiple players have 2-pair? If yes, the player with the highest pair wins
//the pot. If multiple players tie for the highest pair, the player with the
//second highest pair wins the pot. If multiple players tie for both pairs,
//the player with the highest kicker wins the pot. If multiple players tie
//for both pairs and the kicker, the players split the pot.
func TwoPairScore(cards []Card) int {
	prevRank := 0
	finalRank := 0
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 2 {
			if prevRank == 0 {
				prevRank = c.Rank
			} else if prevRank != 0 && prevRank != c.Rank {
				finalRank = c.Rank
			}
		}
	}
	if finalRank == 0 {
		return 0
	}
}

func PairScore(cards []Card) bool {
	for _, c := range cards {
		if len(GetRanks(cards, c.Rank)) == 2 {
			return true
		}
	}
	return false
}
