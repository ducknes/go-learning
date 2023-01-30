package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// HomeWork: Без примитивов синхронизации (без пакета sync) написать калькулятор на горутинах.
// 2+2,3+6,7*7,9/3
// 4, 9, 49, 3

func main() {
	ch := make(chan int)
	strExp := "2 + 2,3 + 6,7 * 7,9 / 3,32 * 9"
	go Calc(StrToSlice(strExp), ch)
	for i := range ch {
		fmt.Print(i, " ")
	}
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
		temp := strings.Split(expressions[i], " ")
		switch temp[1] {
		case "+":
			num1, _ := strconv.Atoi(temp[0])
			num2, _ := strconv.Atoi(temp[2])
			go Add(num1, num2, exitChan)
		case "-":
			num1, _ := strconv.Atoi(temp[0])
			num2, _ := strconv.Atoi(temp[2])
			go Take(num1, num2, exitChan)
		case "*":
			num1, _ := strconv.Atoi(temp[0])
			num2, _ := strconv.Atoi(temp[2])
			go Multiple(num1, num2, exitChan)
		case "/":
			num1, _ := strconv.Atoi(temp[0])
			num2, _ := strconv.Atoi(temp[2])
			go Delete(num1, num2, exitChan)
		}
	}
	time.Sleep(1 * time.Millisecond)
	close(exitChan)
}
