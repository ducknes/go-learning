package main

import (
	"fmt"
	"time"
)

// Горутины

// func main() {
// 	ch := make(chan int)
// 	go sayHello(ch)

// 	for i := range ch {
// 		fmt.Println(i)
// 	}
// }

// func say(word string) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(word)
// }

// func sayHello(exit chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		exit <- i
// 	}

// 	close(exit)
// }

func main() {
	data := make(chan int)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// HomeWoke: Без примитивов синхронизации (без пакета sync) написать калькулятор на горутинах.
// 2 + 2, 3 + 6, 7 * 7, 9 / 3
// 4, 9, 49, 3
