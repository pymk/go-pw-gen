package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

// main is the entry point for the program. It optionally accepts a single
// command-line argument which should be an integer. If the argument is provided,
// the program prints the provided amount of randomly generated passwords. If no
// argument is provided, the program defaults to 1.
func main() {
	num := 1
	var err error

	if len(os.Args) == 2 {
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("The argument must be an integer.")
			return
		}
	}

	for i := 0; i < num; i++ {
		fmt.Println(generator())
	}
}

// generator creates a random string by generating a combination of uppercase
// alphabetic characters, lowercase alphabetic characters, special characters,
// and digits. It creates 6 sets of these characters, shuffles them to ensure
// a random order, and then concatenates them into a single string which it returns.
func generator() string {
	a := alphabet()
	s := special()
	d := digits()
	c := []string{}

	for i := 0; i < 6; i++ {
		c1 := strings.ToUpper((a[rand.IntN(len(a))]))
		c2 := strings.ToLower(a[rand.IntN(len(a))])
		c3 := s[rand.IntN(len(s))]
		c4 := d[rand.IntN(len(d))]
		c = append(c, c1, c2, c3, c4)
	}

	// Shuffle the slice
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})

	return strings.Join(c, "")

}

// alphabet generates a slice of strings containing the uppercase letters
// of the English alphabet (A-Z) by iterating through the ASCII values of
// these characters and converting them to strings.
func alphabet() []string {
	alphabet := make([]string, 26)
	for i := 0; i < 26; i++ {
		// ASCII value of A is 65. So when i is 0, rune('A' + 0) is rune(65).
		// Since characters in the ASCII table are sequential, we can
		// leverages this to create A-Z alphabet.
		alphabet[i] = string(rune('A' + i)) // rune literals in single quotes
	}
	return alphabet
}

// special generates a slice of strings containing a predefined set of special
// characters. It converts a slice of rune literals representing special characters
// into their string equivalents.
func special() []string {
	specialChars := []rune{
		'!', '~', '#', '$', '%', '&', '(', ')', '*',
		'+', ',', '-', '.', ':', ';', '<', '=', '>',
		'?', '@', '[', ']', '^', '_', '{', '|', '}',
	}
	specialCharStrings := make([]string, len(specialChars))
	for i, char := range specialChars {
		specialCharStrings[i] = string(char)
	}
	return specialCharStrings
}

// digits returns a slice of strings containing the digits 0-9. This slice
// represents each digit as a string.
func digits() []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
}
