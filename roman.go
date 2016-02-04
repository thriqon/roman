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

type Number string

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
