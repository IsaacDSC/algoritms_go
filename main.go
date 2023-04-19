package main

import (
	"algoritms/flat"
	"algoritms/reduce"
	"algoritms/removedequals"
	"fmt"
)

var (
	List2d = [][]int{
		{1, 2, 3, 4, 5, 2, 1, 2, 10, 9, 7},
		{6, 7, 8, 9, 10, 2, 3, 4, 6, 7, 8},
	}
)

func main() {
	list1d := flat.Execute(List2d)
	fmt.Println(list1d)
	listNotEquals := removedequals.Execute(list1d)
	fmt.Println(listNotEquals)
	sum := reduce.Execute(listNotEquals, func(acc, current int) int {
		return acc + current
	}, 0)
	fmt.Println(sum)
	divided := reduce.Execute(listNotEquals, func(acc float64, current int) float64 {
		return acc + float64(current)/10.0
	}, 0)
	fmt.Println(divided)
}
