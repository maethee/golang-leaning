package main

import (
	"fmt"
)

func main() {
	ConversionTypeToAnather()
}

// Can create own type
type MyOwn int

// Conversion certain type to anather type
var ConverA int

var ConverB MyOwn

var ConverC string

func ConversionTypeToOther() {
	ConverA = 40
	ConverB = 50

	// we can convert to underlying type
	// ConverC = "test"
	fmt.Printf("%v\t%T\n", ConverA, ConverA)
	fmt.Printf("%v\t%T\n", ConverB, ConverB)
}


//Notes deep in each integer type

//integer type
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32
