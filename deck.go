package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
type deck []string

func newCard() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 13; j++ {
	// 		cards = append(cards, fmt.Sprintf("%s of %s", cardValues[j], cardSuits[i]))
	// 		// cards = append(cards, fmt.Sprintf(cardValues[j], cardSuits[i]))
	// 	}
	// }
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) printFunc() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(card deck, handSize int) (deck, deck) {
	return card[0:handSize], card[handSize:0]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(fileName string) deck {
	byte, error := os.ReadFile(fileName)
	if error != nil {
		fmt.Println("ERROR: ", error)
		os.Exit(1)
	}

	s := strings.Split(string(byte), "\n")
	return deck(s)
}

func (d deck) shuffle() {
	time := time.Now().UnixNano()
	source := rand.NewSource(time)
	r := rand.New(source)
	fmt.Println(time, "\n", r)
	// for i := range d {
	// 	fmt.Println("i: ", i)
	// 	j := rand.Intn(len(d) - 1)
	// 	fmt.Println("j: ", j)
	// 	d[i], d[j] = d[j], d[i]
	// }
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
