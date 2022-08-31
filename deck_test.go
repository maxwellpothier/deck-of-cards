package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length or 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card Two of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected first card Ace of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length or 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
