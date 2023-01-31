package main

import (
	"fmt"
	"io"
	"strings"
)

func main2() {
	r := strings.NewReader("hello world")

	arr := make([]byte, 11)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d, err = %v, b = %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
