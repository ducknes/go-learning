package main

// Functions and methods

import "fmt"

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

// Замыкание
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	// sumCallback := func(n1, n2 int) int {
	// 	return n1 + n2
	// }

	// result := doSomething(sumCallback, "plus")
	// fmt.Println(result)

	// multipleCallback := func(n1, n2 int) int {
	// 	return n1 * n2
	// }

	// resultM := doSomething(multipleCallback, "multiple")
	// fmt.Println(resultM)

	//orderPrice := totalPrice(1)
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))

	p := Point{1, 2}
	//fmt.Println(p)
	//fmt.Println(p.movePoint(1, 1))
	//fmt.Println(p)
	//p.movePointPtr(1, 1)
	//fmt.Println(p)

	v := &p
	v.movePoint(1, 1)
	fmt.Println(p)
	v.movePointPtr(2, 3)
	fmt.Println(v)
}
