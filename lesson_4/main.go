package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Iterator for range and map struct

type Point struct {
	X int `mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("call point method")
}

func main() {
	// arr := []string{"a", "b", "c"}
	// for _, l := range arr {
	// 	fmt.Println(l)
	// }

	pointsMap := map[string]int{
		"x": 123,
		"y": 234,
	}

	for key, value := range pointsMap {
		fmt.Println(key, value)
	}

	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)
}

func mapExample() {
	pointsMap := map[string]Point{
		"b": {13, 15},
	}
	// otherMap := map[string]int{}
	makeMap := make(map[int]Point)

	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	// fmt.Println(pointsMap)
	// fmt.Println(pointsMap["a"])

	makeMap[1] = Point{1, 2}
	// fmt.Println(makeMap)
	// fmt.Println(makeMap[1])

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {1, 2},
		}
	} else {
		oneMorePointsMap["a"] = Point{1, 2}
	}
	fmt.Println(oneMorePointsMap["a"].X)
	fmt.Println(oneMorePointsMap["a"].Y)
	oneMorePointsMap["a"].method()

	key := "v"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key = %s exists in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Printf("key = %s doesn't exists in map\n", key)
		fmt.Println(value)
	}

}
