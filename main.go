package main

import "fmt"

func golearning() {
	fmt.Println("hello world")

	var a int8 = 127              // -128 <> 127, 1 byte, 8 bit
	var b int16 = 32767           // -32768 <> 32767, 2 byte, 16 bit
	var c int32 = 30              // -2bil <> 2bil, 4 byte
	var d int64 = -19000000000000 // -9pent <> 9pent, 8 byte

	var e uint8 = 10 // 0 <> 255, 1 byte
	var f uint16     // 0 <> 65535, 2 byte
	var g uint32     // 0 <> 4bil, 4 byte
	var h uint64     // 0 <> 18pent, 8byte

	var i byte // synonym uint8 0 <> 255
	var j rune // synonim int32
	var k int  //
	var m uint //

	var a1 float32 = 1.8                    // 1.4*10^-45 <> 3.4*10^38, 4 byte
	var b1 float64 = -12222222.888888888888 // 4.8*10^-320 <> 1.8*10^305, 8 byte

	var complex complex64 = 1 + 2i    // float32
	var complex2 complex128 = 4 + 90i // float64

	var bool1 bool = true
	var bool2 bool = false

	var name string = "Ilya"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, m, a1, b1, complex, complex2, bool1, bool2, name)

	///////////////////////////////////////

	var (
		name2 = "Natail"
		age2  = 19
	)

	fmt.Println(name2, age2)

	var name3, age3 = "Elena", 51
	fmt.Println(name3, age3)

	concat := fmt.Sprintf("My name is %s. And I am %d years old", name, age2)
	fmt.Println(concat)

}
