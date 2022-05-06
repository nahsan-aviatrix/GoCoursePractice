package learning

import "testing"

func TestNewDeck(t *testing.T) {
	// check that a deck is created with right num of cards
	// check that first card is correct
	// check if last card is correct
	// OVERALL: checking that we are getting correct output, by checking our assertions about different things.
	// In this case we can check deck size, first card value last card value, etc.

	// mod file responsible for
	// mod

	d := NewDeck()

	if len(d) != 16 {
		// notify t something went wrong:
		t.Errorf("%vExpected deck length of 20, but got:", len(d))
	}
}

// func TestSaveToDeckandNewDeckFromFile() {

// }
