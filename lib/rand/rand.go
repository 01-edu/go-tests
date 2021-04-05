package lib

import (
	"math/big"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1

	StrLen   = 13 // Default length of random strings
	SliceLen = 8  // Default length of slices
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))

	// charsets
	Digit = RuneRange('0', '9')         // Decimal digit characters
	Lower = RuneRange('a', 'z')         // Lowercase latin alphabet characters
	Upper = RuneRange('A', 'Z')         // Uppercase latin alphabet characters
	ASCII = RuneRange(' ', '~')         // ASCII printable characters
	Space = strings.Repeat(" ", StrLen) // Spaces characters
	Basic = Lower + Upper               // Lower and Upper characters
	Alnum = Basic + Digit               // Basic and Digit characters
	Words = Alnum + Space               // Alnum and Space characters
)

func init() {
	rand.Seed(nsSince1970)
}

// RuneRange returns a string containing all the valid runes from a to b.
func RuneRange(a, b rune) string {
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

// IntRange returns a slice containing all the int from a to b.
func IntRange(a, b int) (s []int) {
	for {
		s = append(s, a)
		if a == b {
			return
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}

// RandIntBetween returns a random int between a and b included.
func RandIntBetween(a, b int) int {
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

// RandPosZ returns a random int between 0 and MaxInt included.
func RandPosZ() int { return RandIntBetween(0, MaxInt) }

// RandPos returns a random int between 1 and MaxInt included.
func RandPos() int { return RandIntBetween(1, MaxInt) }

// RandInt returns a random int between MinInt and MaxInt included.
func RandInt() int { return RandIntBetween(MinInt, MaxInt) }

// RandNeg returns a random int between MinInt and 1 included.
func RandNeg() int { return RandIntBetween(MinInt, 1) }

// RandNegZ returns a random int between MinInt and 0 included.
func RandNegZ() int { return RandIntBetween(MinInt, 0) }

// MakeIntFunc returns a slice of ints created by f.
func MakeIntFunc(f func() int) (s []int) {
	i := 0
	for i < SliceLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandPosZ returns a slice of random ints between 0 and MaxInt included.
func MultRandPosZ() []int { return MakeIntFunc(RandPosZ) }

// MultRandPos returns a slice of random ints between 1 and MaxInt included.
func MultRandPos() []int { return MakeIntFunc(RandPos) }

// MultRandInt returns a slice of random ints between MinInt and MaxInt included.
func MultRandInt() []int { return MakeIntFunc(RandInt) }

// MultRandNeg returns a slice of random ints between MinInt and 1 included.
func MultRandNeg() []int { return MakeIntFunc(RandNeg) }

// MultRandNegZ returns a slice of random ints between MinInt and 0 included.
func MultRandNegZ() []int { return MakeIntFunc(RandNegZ) }

// MultRandIntBetween returns a slice of random ints between a and b included.
func MultRandIntBetween(a, b int) []int {
	return MakeIntFunc(func() int { return RandIntBetween(a, b) })
}

// RandRune returns a random printable rune
// (although you may not have the corresponding glyph).
// One-in-ten chance to get a rune higher than 0x10000 (1<<16).
func RandRune() rune {
	ranges := unicode.PrintRanges
	table := ranges[rand.Intn(len(ranges))]
	if rand.Intn(10) == 0 {
		r := table.R32[rand.Intn(len(table.R32))]
		n := uint32(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
		return rune(r.Lo + n*r.Stride)
	}
	r := table.R16[rand.Intn(len(table.R16))]
	n := uint16(rand.Intn(int((r.Hi-r.Lo)/r.Stride) + 1))
	return rune(r.Lo + n*r.Stride)
}

// RandStr returns a string with l random characters taken from chars.
// If chars is empty, the characters are random printable runes.
func RandStr(l int, chars string) string {
	if l <= 0 {
		return ""
	}
	dst := make([]rune, l)
	if chars == "" {
		for i := range dst {
			dst[i] = RandRune()
		}
	} else {
		src := []rune(chars)
		for i := range dst {
			r := rand.Intn(len(src))
			dst[i] = src[r]
		}
	}
	return string(dst)
}

// RandDigit returns a string containing random decimal digit characters.
func RandDigit() string { return RandStr(StrLen, Digit) }

// RandLower returns a string containing random lowercase latin alphabet characters.
func RandLower() string { return RandStr(StrLen, Lower) }

// RandUpper returns a string containing random uppercase latin alphabet characters.
func RandUpper() string { return RandStr(StrLen, Upper) }

// RandASCII returns a string containing random ASCII printable characters.
func RandASCII() string { return RandStr(StrLen, ASCII) }

// RandSpace returns a string containing random spaces characters.
func RandSpace() string { return RandStr(StrLen, Space) }

// RandBasic returns a string containing random lower and upper characters.
func RandBasic() string { return RandStr(StrLen, Basic) }

// RandAlnum returns a string containing random basic and digit characters.
func RandAlnum() string { return RandStr(StrLen, Alnum) }

// RandWords returns a string containing random alphanumeric and space characters.
func RandWords() string { return RandStr(StrLen, Words) }

// MakeStrFunc returns a slice of strings created by f.
func MakeStrFunc(f func() string) (s []string) {
	i := 0
	for i < StrLen {
		s = append(s, f())
		i++
	}
	return
}

// MultRandDigit returns a slice of strings containing random Decimal digit characters.
func MultRandDigit() []string { return MakeStrFunc(RandDigit) }

// MultRandLower returns a slice of strings containing random Lowercase latin alphabet.
func MultRandLower() []string { return MakeStrFunc(RandLower) }

// MultRandUpper returns a slice of strings containing random Uppercase latin alphabet.
func MultRandUpper() []string { return MakeStrFunc(RandUpper) }

// MultRandASCII returns a slice of strings containing random ASCII printable characters.
func MultRandASCII() []string { return MakeStrFunc(RandASCII) }

// MultRandSpace returns a slice of strings containing random Spaces characters.
func MultRandSpace() []string { return MakeStrFunc(RandSpace) }

// MultRandBasic returns a slice of strings containing random Lower and Upper characters.
func MultRandBasic() []string { return MakeStrFunc(RandBasic) }

// MultRandAlnum returns a slice of strings containing random Basic and Digit characters.
func MultRandAlnum() []string { return MakeStrFunc(RandAlnum) }

// MultRandWords returns a slice of strings containing random Alnum and Space characters.
func MultRandWords() []string { return MakeStrFunc(RandWords) }
