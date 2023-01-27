package main

import (
	"errors"
	"fmt"
	"log"
)

// Panics!

// type name struct {
// 	A, B int
// }

// func (n *name) method() {
// 	fmt.Println("ok")
// 	fmt.Println(n.A)
// }

type AppError struct {
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Message
}

func main() {
	// n := &name{1, 2}
	// n.method()

	divide(4, 0)
	fmt.Println("after panic")
}

func divide(a, b int) {
	// обработка паники
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("this is AppError handle panic!", err)
				} else {
					fmt.Println("this is Other Error handle panic!", err)
				}
			default:
				panic("this is default go panic")
			}
			log.Println("panic happend:", err)
		}
	}()
	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		//panic(fmt.Errorf("some error"))

		panic(&AppError{
			Message: "this in divide by zero custom error",
			Err:     nil,
		})

		//panic("error")
	}
	return a / b
}
