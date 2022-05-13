package main

import (
	"log"
	"testing"
)

func TestIsLetter(t *testing.T) {
	notLetter := isLetter("ap4le")
	isLetter := isLetter("fjsdfj")

	if notLetter != false {
		log.Fatalf("Expected false, got %t", notLetter)
	}

	if isLetter != true {
		log.Fatalf("Expected true, got %t", isLetter)
	}
}
