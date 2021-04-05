package random

import (
	"math/big"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1

	strLen   = 13 // Default length of random strings
	sliceLen = 8  // Default length of slices
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))

	// charsets
	Digit = runeRange('0', '9')         // Decimal digit characters
	Lower = runeRange('a', 'z')         // Lowercase latin alphabet characters
	Upper = runeRange('A', 'Z')         // Uppercase latin alphabet characters
	ASCII = runeRange(' ', '~')         // ASCII printable characters
	Space = strings.Repeat(" ", strLen) // Spaces characters
	Basic = Lower + Upper               // Lower and Upper characters
	Alnum = Basic + Digit               // Basic and Digit characters
	Words = Alnum + Space               // Alnum and Space characters
)

func init() {
	rand.Seed(nsSince1970)
}

func runeRange(a, b rune) string {
	var s []rune
	for {
		if utf8.ValidRune(a) {
			s = append(s, a)
		}
		if a == b {
			return string(s)
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

func makeIntFunc(f func() int) (s []int) {
	for i := 0; i < sliceLen; i++ {
		s = append(s, f())
	}
	return
}

func str(chars string) string {
	runes := []rune(chars)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

var Intn = rand.Intn

// IntBetween returns a random int between a and b included.
func IntBetween(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := big.NewInt(int64(b))      // b
	n.Sub(n, big.NewInt(int64(a))) // b-a
	n.Add(n, big.NewInt(1))        // b-a+1
	n.Rand(bigRand, n)             // 0 <= n <= b-a
	n.Add(n, big.NewInt(int64(a))) // a <= n <= b
	return int(n.Int64())
}

// Int returns a random int between MinInt and MaxInt included.
func Int() int {
	return IntBetween(MinInt, MaxInt)
}

// Ints returns a slice of random ints between MinInt and MaxInt included.
func Ints() []int {
	return makeIntFunc(Int)
}

// IntsBetween returns a slice of random ints between a and b included.
func IntsBetween(a, b int) []int {
	return makeIntFunc(func() int {
		return IntBetween(a, b)
	})
}

// RandStr returns a string with l random characters taken from chars.
func RandStr(length int, chars string) (dst string) {
	if length <= 0 {
		return ""
	}
	if chars == "" {
		panic("No charset provided")
	}
	for ; length > 0; length-- {
		r := rand.Intn(len(chars))
		dst += string(chars[r])
	}
	return string(dst)
}

// RandDigit returns a string containing random decimal digit characters.
func RandDigit() string { return str(Digit) }

// RandLower returns a string containing random lowercase latin alphabet characters.
func RandLower() string { return str(Lower) }

// RandUpper returns a string containing random uppercase latin alphabet characters.
func RandUpper() string { return str(Upper) }

// RandASCII returns a string containing random ASCII printable characters.
func RandASCII() string { return str(ASCII) }

// RandSpace returns a string containing random spaces characters.
func RandSpace() string { return str(Space) }

// RandBasic returns a string containing random lower and upper characters.
func RandBasic() string { return str(Basic) }

// RandAlnum returns a string containing random basic and digit characters.
func RandAlnum() string { return str(Alnum) }

// RandWords returns a string containing random alphanumeric and space characters.
func RandWords() string { return str(Words) }

func makeStrFunc(f func() string) (s []string) {
	for i := 0; i < strLen; i++ {
		s = append(s, f())
	}
	return
}

// MultRandDigit returns a slice of strings containing random Decimal digit characters.
func MultRandDigit() []string { return makeStrFunc(RandDigit) }

// MultRandLower returns a slice of strings containing random Lowercase latin alphabet.
func MultRandLower() []string { return makeStrFunc(RandLower) }

// MultRandUpper returns a slice of strings containing random Uppercase latin alphabet.
func MultRandUpper() []string { return makeStrFunc(RandUpper) }

// MultRandASCII returns a slice of strings containing random ASCII printable characters.
func MultRandASCII() []string { return makeStrFunc(RandASCII) }

// MultRandSpace returns a slice of strings containing random Spaces characters.
func MultRandSpace() []string { return makeStrFunc(RandSpace) }

// MultRandBasic returns a slice of strings containing random Lower and Upper characters.
func MultRandBasic() []string { return makeStrFunc(RandBasic) }

// MultRandAlnum returns a slice of strings containing random Basic and Digit characters.
func MultRandAlnum() []string { return makeStrFunc(RandAlnum) }

// MultRandWords returns a slice of strings containing random Alnum and Space characters.
func MultRandWords() []string { return makeStrFunc(RandWords) }
