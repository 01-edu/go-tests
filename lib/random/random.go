package random

import (
	"math/big"
	"math/rand"
	"time"
)

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))
)

func init() {
	rand.Seed(nsSince1970)
}

func makeIntFunc(f func() int) (s []int) {
	for i := 0; i < 8; i++ {
		s = append(s, f())
	}
	return
}

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

// IntSlice returns a slice of 8 random ints between MinInt and MaxInt included.
func IntSlice() []int {
	return makeIntFunc(Int)
}

// IntSliceBetween returns a slice of 8 random ints between a and b included.
func IntSliceBetween(a, b int) []int {
	return makeIntFunc(func() int {
		return IntBetween(a, b)
	})
}

// Str returns a string with l random characters taken from chars.
func Str(chars string, length int) (dst string) {
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

// Str returns a slice of 8 strings with 13 random characters taken from chars.
func StrSlice(chars string) (s []string) {
	for i := 0; i < 8; i++ {
		s = append(s, Str(chars, 13))
	}
	return
}
