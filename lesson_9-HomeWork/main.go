package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "2 + 2,3 + 6,7 * 7,9 / 3"

	s := strings.Split(str, ",")
	fmt.Println(s)

	var spl [][]string
	for i := 0; i < len(s); i++ {
		spl = append(spl, strings.Split(s[i], " "))
	}
	fmt.Println(spl) 
}
