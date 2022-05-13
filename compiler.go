package main

import (
	"fmt"
	"unicode"
)

type token struct {
	kind  string
	value string
}

// isLetter checks if the character passed in is a letter. Returns true if it is a letter, false otherwise
func isLetter(char string) bool {
	for _, k := range char {
		if !unicode.IsLetter(k) {
			return false
		}
	}
	return true
}

// tokenizer takes a python function as input and tokenizes it
func tokenizer(input string) {
	input += "\n"

	tokens := []token{}

	for i := 0; i < len([]rune(input)); i++ {
		// get the current character
		char := string(input[i])

		// find all the word tokens
		if isLetter(char) {
			letters := ""
			// iterate over all the letters
			for isLetter(char) {
				letters += char
				i++
				char = string(input[i])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: letters,
			})
		}
	}
	fmt.Println(fmt.Sprintf("tokenizer tokens: %v", tokens)) // __AUTO_GENERATED_PRINT_VAR__
}

func main() {
	tokenizer("def hello()")
}
