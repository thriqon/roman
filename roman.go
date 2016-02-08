// Package roman provides fast and simple conversion methods for Roman
// numerals.
package roman

import "strings"

var romanPrefixes = []struct {
	prefix string
	val    uint
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// A Number is a romanly encoded unsigned integer. Internally, it's based on a
// string, so all common string operations are available. A Number can be either constructed by
// direct instantiation (only safe if the user is sure that it will parse, e.g. for constants)
// with roman.Number("XXI"), or by converting from a native uint value with FromInt.
type Number string

// AsInt gives the value of the romanly encoded numeral as uint. It panics if any of the characters
// in the Number are not valid at this point (e.g. unknown characters such as 'A').
func (r Number) AsInt() uint {
	var sum uint

outer:
	for r != "" {
		for _, el := range romanPrefixes {
			if strings.HasPrefix(string(r), el.prefix) {
				r = r[len(el.prefix):]
				sum += el.val
				continue outer
			}
		}
		panic("invalid character in sequence: " + r)
	}
	return sum
}

// Add adds the given number to the receiver, and returns the result as roman
// numeral.  Internally, this converts both numbers to uint, adds them, and
// then encodes it back to roman.
func (r Number) Add(x Number) Number {
	return FromInt(r.AsInt() + x.AsInt())
}

// FromInt encodes the integer passed as parameter as roman numeral.
func FromInt(i uint) Number {
	var res string
	var prefixPos int

	for i > 0 {
		p := romanPrefixes[prefixPos]
		for i >= p.val {
			i -= p.val
			res += p.prefix
		}
		prefixPos++
	}
	return Number(res)
}
