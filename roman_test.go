package roman

import (
	"fmt"
	"testing"
)

var cases = []struct {
	x Number
	i uint
}{
	{"I", 1},
	{"II", 2},
	{"III", 3},
	{"IV", 4},
	{"V", 5},
	{"VI", 6},
	{"VII", 7},
	{"VIII", 8},
	{"IX", 9},
	{"X", 10},
	{"XI", 11},
	{"XII", 12},
	{"XXIX", 29},
	{"XXXIX", 39},
	{"XL", 40},
	{"L", 50},
	{"XC", 90},
	{"XCII", 92},
	{"C", 100},
	{"D", 500},
	{"CD", 400},
	{"M", 1000},
	{"MM", 2000},
	{"MCMLXXXVII", 1987},
	{"DCC", 700},
}

func TestAsInt(t *testing.T) {
	for i, el := range cases {
		if actual := el.x.AsInt(); actual != el.i {
			t.Errorf("Case %v failed: %v should be %v, but was %v", i, el.x, el.i, actual)
		}
	}
}

func TestAsIntPanics(t *testing.T) {
	defer func() {
		x := recover()
		if x == nil {
			t.Error("Did not panic")
		}
	}()

	Number("a").AsInt()
}

func TestAddition(t *testing.T) {
	cases := []struct {
		a, b Number
		res  Number
	}{
		{"I", "I", "II"},
		{"I", "II", "III"},
		{"III", "III", "VI"},
		{"MMII", "II", "MMIV"},
	}

	for i, el := range cases {
		if actual := el.a.Add(el.b); actual != el.res {
			t.Errorf("Case %v failed: %v + %v should be %v, but was %v", i, el.a, el.b, el.res, actual)
		}
	}
}

func TestFromInt(t *testing.T) {
	for i, el := range cases {
		if actual := FromInt(el.i); actual != el.x {
			t.Errorf("Case %v failed: %v should be %v, but was %v", i, el.i, el.x, actual)
		}
	}
}

func ExampleFromInt() {
	fmt.Println(FromInt(1))
	fmt.Println(FromInt(1999))
	// Output:
	// I
	// MCMXCIX
}
