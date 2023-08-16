package random

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"reflect"
	"time"
)

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))

	MinSliceLen = 1
	MaxSliceLen = 8

	StringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	MinStringLen  = 0
	MaxStringLen  = 64
)

func init() {
	rand.Seed(nsSince1970)
}

func makeIntFunc(f func() int) (s []int) {
	for i := MinSliceLen; i < MaxSliceLen; i++ {
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
func Str() (dst string) {
	for length := 0; length < IntBetween(MinStringLen, MaxStringLen); length++ {
		r := rand.Intn(len(StringCharset))
		dst += string(StringCharset[r])
	}
	return string(dst)
}

// Str returns a slice of n strings of random sizes with random characters from StringCharset.
func StrSlice(n int) (s []string) {
	for i := MinSliceLen; i < MaxSliceLen; i++ {
		s = append(s, Str())
	}
	return
}

func GenerateValue(i interface{}) interface{} {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		return Int()
	case reflect.String:
		return Str()
	case reflect.Slice:
		//TODO:
		return nil
	default:
		fmt.Fprintf(os.Stderr, "Can't create random test arguments for type: %v, %v\n", reflect.TypeOf(i).Kind(), i)
	}
	return nil
}
