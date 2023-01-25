package main

import "fmt"

// Interfaces

type structHere struct {
	N1, N2 int
}

func (s *structHere) Sum() int {
	return s.N1 + s.N2
}

type InterfaceHere interface {
	Sum() int
}

func main() {
	//var a InterfaceHere
	// sh := structHere{1, 2}
	// os := otherStruct{2, 3}

	// a = &sh
	// fmt.Println(a.Sum())
	// a = os
	// fmt.Println(a.Sum())

	// var os *structHere
	// var i InterfaceHere
	// i = os
	// fmt.Println(i.Sum())
	// fmt.Printf("(%v, %T)\n", i, i)
	// if i == nil {
	// 	fmt.Println("nil")
	// } else {
	// 	fmt.Println("not nil")
	// }

	var a interface{} = "hello"

	// s := a.(string)
	// fmt.Println(s)

	// s, ok := a.(string)
	// fmt.Println(s, ok)

	// f, ok := a.(float32)
	// if ok {
	// 	// use f
	// 	fmt.Println(f, ok)
	// }

	switch a.(type) {
	case int:
		fmt.Println()
	}
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
