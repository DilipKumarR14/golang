package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Card struct {
	Type string
	Suit string
}

type Deck []Card

func New() (deck Deck) {

	types := []string{"2", "3", "4", "5", "6", "7",
		"8", "9", "10", "Jack", "Queen", "King", "Ace"}

	// Valid suits include Heart, Diamond, Club & Spade
	suits := []string{"Heart", "Diamond", "Club", "Spade"}

	// Loop over each type and suit appending to the deck
	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

// Shuffle the deck
func Shuffle(d Deck) Deck {
	for i := 1; i < len(d); i++ {
		// Create a random int up to the number of cards
		r := rand.Intn(i + 1)

		// If the the current card doesn't match the random
		// int we generated then we'll switch them out
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}

// Deal a specified amount of cards
func Deal(d Deck, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(d[i])
	}
}

// Debug helps debugging the deck of cards
func Debug(d Deck) {
	if os.Getenv("DEBUG") != "" {
		for i := 0; i < len(d); i++ {
			fmt.Printf("Card #%d is a %s of %ss\n", i+1, d[i].Type, d[i].Suit)
		}
	}
}

// Seed our randomness with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	deck := New()
	Debug(deck)
	Shuffle(deck)
	Debug(deck)
	Deal(deck, 52)
}