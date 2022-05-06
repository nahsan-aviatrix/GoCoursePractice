package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

// can think that this new deck type extends []string
type deck []string

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func (cards deck) Print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entierly quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := (strings.Split(string(bs), ","))
	return deck(s)
}

func (d deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		newPosition := random.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
