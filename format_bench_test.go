package roman

import "testing"

import roman "github.com/StefanSchroeder/Golang-Roman"

func benchmarkFormatThis(x uint, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if r := FromInt(x); r == "" {
			panic("will not happen")
		}
	}
}

func benchmarkFormatSchroeder(x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if r := roman.Roman(x); r == "" {
			panic("will not happen")
		}
	}
}

func benchmarkFormatRosetta(x int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if r, _ := formatRoman(x); r == "" {
			panic("will not happen")
		}
	}
}

// 1

func BenchmarkFormat1This(b *testing.B) {
	benchmarkFormatThis(1, b)
}
func BenchmarkFormat1Schroeder(b *testing.B) {
	benchmarkFormatSchroeder(1, b)
}
func BenchmarkFormat1Rosetta(b *testing.B) {
	benchmarkFormatRosetta(1, b)
}

// 2016

func BenchmarkFormat2016This(b *testing.B) {
	benchmarkFormatThis(2016, b)
}
func BenchmarkFormat2016Schroeder(b *testing.B) {
	benchmarkFormatSchroeder(2016, b)
}
func BenchmarkFormat2016Rosetta(b *testing.B) {
	benchmarkFormatRosetta(2016, b)
}

// 1999

func BenchmarkFormat1999This(b *testing.B) {
	benchmarkFormatThis(1999, b)
}
func BenchmarkFormat1999Schroeder(b *testing.B) {
	benchmarkFormatSchroeder(1999, b)
}
func BenchmarkFormat1999Rosetta(b *testing.B) {
	benchmarkFormatRosetta(1999, b)
}
