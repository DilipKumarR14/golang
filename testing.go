package main

 import (
         "fmt"
         "math/rand"
         "time"
 )

 // let just stick to arrays instead of slice
 // because 4 x 13 is pretty standard for most decks....unless you want to include jokers.

 var suit = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}

 var face = [13]string{"Ace", "Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

 const rows = 13
 const columns = 4
 const total = rows * columns

 // shuffle deck of cards
 func shuffleDeck(workDeck []string) []string {

         shuffled := make([]string, len(workDeck))

         rand.Seed(time.Now().UTC().UnixNano())
         perm := rand.Perm(len(workDeck))

         for i, v := range perm {
                 shuffled[v] = workDeck[i]
         }

         return shuffled
 }

 func main() {

         var initDeck []string

         //initialize or populate our initial deck of cards
         //in sorted form
         for r := 0; r <= rows-1; r++ {
                 for c := 0; c <= columns-1; c++ {
                         tmp := face[r] + " of " + suit[c]
                         initDeck = append(initDeck, tmp)

                 }
         }

         // our new deck of cards is sorted !
         for k, v := range initDeck {
                 fmt.Printf("%d %s\n", k, v)
         }

         fmt.Println("-----shuffling deck------")
         shuffledDeck := shuffleDeck(initDeck)

         // our new deck of cards is now shuffled !
         for k, v := range shuffledDeck {
                 fmt.Printf("%d %s\n", k, v)
         }

 }