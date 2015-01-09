package main

import "math/rand"

type Deck []Card

func NewDeck() (d Deck) {
	for s := Clubs; s <= Spades; s++ {
		for r := Two; r <= Ace; r++ {
			c := Card{r, s}
			d = append(d, c)
		}
	}
	//fmt.Println(d)
	// d.Shuffle()
	return
}

func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Fischer-Yates http://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func (d Deck) Shuffle() {
	for i := len(d) - 1; i >= 1; i-- {
		j := rand.Intn(i)
		d.Swap(i, j)
	}
}

func (d Deck) String() (out string) {
	for _, c := range d {
		out += c.String() + " "
	}
	return
}

func Draw(d Deck) (Deck, Card) {
	i := rand.Intn(len(d))
	c := d[i]
	d = append(d[:i], d[i+1:]...)
	return d, c
}
