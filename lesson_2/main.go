package main

import "fmt"

// ------- Функции, возвращаемые значения --------

const widthTable int = 200

// func main() {
// 	s, s2, s3 := test()
// 	s4 := test1()
// 	fmt.Println(s, s2, s3, s4)
// }

func test() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return // a, b, c
}

func test1() (a string) {
	a = "ok"
	return // "Ilya"
}

// ------- Циклы, ветвления --------

// func main() {
// 	sum := 0

// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}

// 	for sum < 1000 {
// 		sum += 10
// 	}

// 	for sum < 10000 { // while
// 		sum += 1000
// 	}

// 	// for { unfinity loop
// 	// 	sum += 10
// 	// }

// 	a := 0
// 	for a < 1000 {
// 		if a == 100 {
// 			// fmt.Println("a is 100")
// 		} else {
// 			// fmt.Println(fmt.Sprintf("a is not 100, a = %d", a))
// 		}
// 		a++
// 	}

// 	x := 0
// 	for x < 1000 {
// 		// if i := isTest(x); i == 1 {
// 		// 	fmt.Println("x = 2")
// 		// } else if i == 2 {
// 		// 	fmt.Println("x = 3")
// 		// } else {
// 		// 	fmt.Printf("unknown x, x=%d\n", x)
// 		// }
// 		// i := isTest(x)
// 		switch i := isTest(x); i {
// 		case 1:
// 			fmt.Println("x = 2")
// 		case 2:
// 			fmt.Println("x = 3")
// 		default:
// 			fmt.Printf("unknown x, x=%d\n", x)
// 		}
// 		x++
// 	}
// }

// func isTest(a int) int {
// 	if a == 2 {
// 		return 1
// 	} else if a == 3 {
// 		return 2
// 	}
// 	return 3
// }

// ------- Defer --------

func main() {
	// stack, LIFO (Last In First Out)
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("hello ")

}
