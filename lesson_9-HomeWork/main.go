package main

import (
	"fmt"
	"strconv"
	"strings"
)

// HomeWork: Без примитивов синхронизации (без пакета sync) написать калькулятор на горутинах.
// 2+2,3+6,7*7,9/3
// 4, 9, 49, 3

func main() {
	ch := make(chan int)
	// go Add(2, 2, ch)
	// go Add(3, 6, ch)
	// go Multiple(7, 7, ch)
	// go Delete(9, 3, ch)
	strExp := "2+2,3+6,7*7,9/3"
	go Calc(StrToSlice(strExp), ch)
	// for i := range ch {
	// 	fmt.Print(i, " ")
	// 	count++
	// 	fmt.Println(count)
	// }

	fmt.Print(" ", <-ch)
	fmt.Print(" ", <-ch)
	fmt.Print(" ", <-ch)
	fmt.Print(" ", <-ch)
}

func Add(a, b int, exit chan int) {
	exit <- a + b
}

func Take(a, b int, exit chan int) {
	exit <- a - b
}

func Multiple(a, b int, exit chan int) {
	exit <- a * b
}

func Delete(a, b int, exit chan int) {
	exit <- a / b
}

func StrToSlice(s string) []string {
	return strings.Split(s, ",")
}

func Calc(expressions []string, exitChan chan int) {
	for i := range expressions {
		switch expressions[i][1:2] {
		case "+":
			num1, _ := strconv.Atoi(expressions[i][0:1])
			num2, _ := strconv.Atoi(expressions[i][2:])
			exitChan <- num1 + num2
		case "-":
			num1, _ := strconv.Atoi(expressions[i][0:1])
			num2, _ := strconv.Atoi(expressions[i][2:])
			exitChan <- num1 - num2
		case "*":
			num1, _ := strconv.Atoi(expressions[i][0:1])
			num2, _ := strconv.Atoi(expressions[i][2:])
			exitChan <- num1 * num2
		case "/":
			num1, _ := strconv.Atoi(expressions[i][0:1])
			num2, _ := strconv.Atoi(expressions[i][2:])
			exitChan <- num1 / num2
		}
	}
}
