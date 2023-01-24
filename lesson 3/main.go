package main

import "fmt"

// -------Pointers, Stucts, Arrays and Slices--------

type Point struct {
	X int
	Y int
	S string
}

func (p *Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {
	//pointers()

	// p1 := Point{
	// 	X: 1,
	// 	Y: 2,
	// 	S: "Ilya",
	// }
	// p1.method()
	// p2 := &p1
	// p2.method()

	//slices()

	//arraysAndSlices()

	//slicesMistakes()

	// var snil []int
	// fmt.Println(snil, len(snil), cap(snil))
	// if snil == nil {
	// 	fmt.Println("slice is nil!")
	// }

	
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "A",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{
		X: 123,
	}
	fmt.Println(p2)

	g := &p1
	fmt.Println(g)
	fmt.Println(*g)
	fmt.Println(&g)
}

func pointers() {
	a := "hello world"

	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oh my"
	fmt.Println(a)

	b := 42
	g := &b
	*g /= 2
	fmt.Println(b)
}

func arrays() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[1])
	numbers := [...]int{1, 2, 3}
	numbers[2] = 4
}

func slices() {
	// slices
	letters := []string{"a", "b", "c"}
	letters[0] = "1"
	letters = append(letters, "d")
	letters = append(letters, "e", "f", "g")
	// fmt.Println(letters)

	createSlice := make([]string, 3)
	fmt.Printf("len: %d\n", len(createSlice))
	fmt.Printf("cap: %d\n", cap(createSlice))
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	//createSlice[3] = "4" // panic: runtime error: index out of range [3] with length 3
	createSlice = append(createSlice, "4")
	fmt.Println(createSlice)
	fmt.Printf("len: %d\n", len(createSlice))
	fmt.Printf("cap: %d\n", cap(createSlice))
}

func arraysAndSlices() {
	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	a := animalsArr[0:2]
	fmt.Println(a)
	b := animalsArr[1:3]
	fmt.Println(b)

	b[0] = "123"
	fmt.Println(a)
	fmt.Println(animalsArr)
}

func slicesMistakes() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[:5]
	fmt.Println(t)
	tt := s[5:]
	fmt.Println(tt)
	ttt := s[0:]
	fmt.Println(ttt)
	tttt := s[:]
	fmt.Println(tttt)
}

