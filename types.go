package main

import "fmt"

var a bool = true

var b string = "Hello World!"

var c int = 1
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

var d byte; // alias for uint8

var e rune; // alias for int32, represents a Unicode code point

var f float32 = 2.2 // float64

var g complex64 = -5 + 12i //complex128

func main() {
	fmt.Println(a, b, c, d, e, f, g) //true Hello World! 1 0 0 2.2 (-5+12i)
}