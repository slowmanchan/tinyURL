package encoder

import (
	"strings"
)

func Encode(num int) string {
	// turn id into a base 62 to "shrink it"
	digits := []int{}
	for num > 0 {
		remainder := num % 62
		digits = append(digits, remainder)
		num /= 62
	}

	// reverse it (since we started from right to left)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// all your doing now is
	// matching each base 62 digit
	// with an letter from your alphabet set
	alphaArr := strings.Split(Alphabet, "")
	res := []string{}
	for _, d := range digits {
		res = append(res, alphaArr[d])
	}

	return strings.Join(res, "")
}
