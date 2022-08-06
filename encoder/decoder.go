package encoder

import (
	"strings"
)

func Decode(url string) int {
	resInt := 0
	if url == "" {
		return resInt
	}

	// Multiply by 62 each time and add the index of the letter.
	// This effectively reverts the number back to base10
	// its like wif we wanted the "tens place" in a base10 number. We would keep mutliplying by 10
	// to get this next one.
	for _, c := range url {
		resInt = resInt*62 + strings.Index(Alphabet, string(c))
	}

	return resInt
}
