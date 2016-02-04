package roman

import "testing"

import roman "github.com/StefanSchroeder/Golang-Roman"

func benchmarkThis(rn string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if x := Number(rn).AsInt(); x < 0 {
			panic("will not happen")
		}
	}
}

func benchmarkSchroeder(rn string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if x := roman.Arabic(rn); x < 0 {
			panic("will not happen")
		}
	}
}

func benchmarkRosetta(rn string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if x, _ := parseRomanRosetta(rn); x < 0 {
			panic("will not happen")
		}
	}
}
func benchmarkRosettaSimple(rn string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if x := parseRomanRosettaSimple(rn); x < 0 {
			panic("will not happen")
		}
	}
}

// MMCDLXXVIII

func BenchmarkParsingMMCDLXXVIIIThisPackage(b *testing.B) {
	benchmarkThis("MMCDLXXVIII", b)
}

func BenchmarkParsingMMCDLXXVIIISchroeder(b *testing.B) {
	benchmarkSchroeder("MMCDLXXVIII", b)

}

func BenchmarkParsingMMCDLXXVIIIRosetta(b *testing.B) {
	benchmarkRosetta("MMCDLXXVIII", b)
}

func BenchmarkParsingMMCDLXXVIIIRosettaSimple(b *testing.B) {
	benchmarkRosettaSimple("MMCDLXXVIII", b)
}

// I

func BenchmarkParsingIThisPackage(b *testing.B) {
	benchmarkThis("I", b)
}

func BenchmarkParsingISchroeder(b *testing.B) {
	benchmarkSchroeder("I", b)

}

func BenchmarkParsingIRosetta(b *testing.B) {
	benchmarkRosetta("I", b)
}

func BenchmarkParsingIRosettaSimple(b *testing.B) {
	benchmarkRosettaSimple("I", b)
}

// MMXXVI

func BenchmarkParsingMMXXVIThisPackage(b *testing.B) {
	benchmarkThis("MMXXVI", b)
}

func BenchmarkParsingMMXXVISchroeder(b *testing.B) {
	benchmarkSchroeder("MMXXVI", b)

}

func BenchmarkParsingMMXXVIRosetta(b *testing.B) {
	benchmarkRosetta("MMXXVI", b)
}

func BenchmarkParsingMMXXVIRosettaSimple(b *testing.B) {
	benchmarkRosettaSimple("MMXXVI", b)
}

// MCMXCIX

func BenchmarkParsingMCMXCIXThisPackage(b *testing.B) {
	benchmarkThis("MCMXCIX", b)
}

func BenchmarkParsingMCMXCIXSchroeder(b *testing.B) {
	benchmarkSchroeder("MCMXCIX", b)

}

func BenchmarkParsingMCMXCIXRosetta(b *testing.B) {
	benchmarkRosetta("MCMXCIX", b)
}

func BenchmarkParsingMCMXCIXRosettaSimple(b *testing.B) {
	benchmarkRosettaSimple("MCMXCIX", b)
}
