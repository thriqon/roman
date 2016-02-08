package benchmarking

import (
	"errors"
	"fmt"
)

// TAKEN FROM Rosetta Code
// Only used for benchmarking

func parseRomanRosetta(s string) (r int, err error) {
	var m = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	if s == "" {
		return 0, errors.New("Empty string")
	}
	is := []rune(s) // easier to convert string up front
	var c0 rune     // c0: roman character last read
	var cv0 int     // cv0: value of cv

	// the key to the algorithm is to process digits from right to left
	for i := len(is) - 1; i >= 0; i-- {
		// read roman digit
		c := is[i]
		k := c == '\u0305' // unicode overbar combining character
		if k {
			if i == 0 {
				return 0, errors.New(
					"Overbar combining character invalid at position 0")
			}
			i--
			c = is[i]
		}
		cv := m[c]
		if cv == 0 {
			if c == 0x0305 {
				return 0, fmt.Errorf(
					"Overbar combining character invalid at position %d", i)
			} else {
				return 0, fmt.Errorf(
					"Character unrecognized as Roman digit: %c", c)
			}
		}
		if k {
			c = -c // convention indicating overbar
			cv *= 1000
		}

		// handle cases of new, same, subtractive, changed, in that order.
		switch {
		default: // case 4: digit change
			fallthrough
		case c0 == 0: // case 1: no previous digit
			c0 = c
			cv0 = cv
		case c == c0: // case 2: same digit
		case cv*5 == cv0 || cv*10 == cv0: // case 3: subtractive
			// handle next digit as new.
			// a subtractive digit doesn't count as a previous digit.
			c0 = 0
			r -= cv  // subtract...
			continue // ...instead of adding
		}
		r += cv // add, in all cases except subtractive
	}
	return r, nil
}

// function, per task description
func parseRomanRosettaSimple(roman string) (arabic int) {
	var m = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last_digit := 1000
	for _, r := range roman {
		digit := m[r]
		if last_digit < digit {
			arabic -= 2 * last_digit
		}
		last_digit = digit
		arabic += digit
	}

	return arabic
}

func formatRoman(n int) (string, bool) {
	var (
		m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		m3 = []string{"", "M", "MM", "MMM", "I̅V̅",
			"V̅", "V̅I̅", "V̅I̅I̅", "V̅I̅I̅I̅", "I̅X̅"}
		m4 = []string{"", "X̅", "X̅X̅", "X̅X̅X̅", "X̅L̅",
			"L̅", "L̅X̅", "L̅X̅X̅", "L̅X̅X̅X̅", "X̅C̅"}
		m5 = []string{"", "C̅", "C̅C̅", "C̅C̅C̅", "C̅D̅",
			"D̅", "D̅C̅", "D̅C̅C̅", "D̅C̅C̅C̅", "C̅M̅"}
		m6 = []string{"", "M̅", "M̅M̅", "M̅M̅M̅"}
	)

	if n < 1 || n >= 4e6 {
		return "", false
	}
	// this is efficient in Go.  the seven operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return m6[n/1e6] + m5[n%1e6/1e5] + m4[n%1e5/1e4] + m3[n%1e4/1e3] +
			m2[n%1e3/1e2] + m1[n%100/10] + m0[n%10],
		true
}
