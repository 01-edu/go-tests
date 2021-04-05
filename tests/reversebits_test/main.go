package main

import (
	"math/bits"
	"reflect"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func reverseBits(octet byte) byte {
	return bits.Reverse8(uint8(octet))
}

func challengeBytes(fn1, fn2 interface{}, args ...interface{}) {
	st1 := challenge.Monitor(fn1, args)
	st2 := challenge.Monitor(fn2, args)
	if !reflect.DeepEqual(st1.Results, st2.Results) {
		challenge.Fatalf("%s(%08b) == %08b instead of %08b\n",
			"ReverseBits",
			args[0].(byte),
			st1.Results[0].(byte),
			st2.Results[0].(byte),
		)
	}
}

func main() {
	args := []byte{0x26, 0x27, 0x28, 0x29, 0xAA, 0xBB}

	for i := 0; i < 10; i++ {
		n := random.IntBetween(0, 255)
		args = append(args, byte(n))
	}

	for _, v := range args {
		challengeBytes(student.ReverseBits, reverseBits, v)
	}
}
