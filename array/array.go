package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5}

var multiarr = [][]int{
	{1, 2, 3, 4, 5},
	{7, 8, 9, 10, 11}}
var emptyArr = make([]int, 0)

func main() {
	for _, y := range arr {
		fmt.Printf("%d \n", y)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d \n", arr[i])
	}
	fmt.Println("Multi Dimensional Array")
	for _, x := range multiarr {
		for _, y := range x {
			fmt.Printf("%d \n", y)
		}
	}
	for i := 0; i < len(multiarr); i++ {
		for j := 0; j < len(multiarr[i]); j++ {
			el := multiarr[i][j]
			fmt.Printf("%d \n", el)
			emptyArr = append(emptyArr, el)
		}
	}
	fmt.Printf("len:%d  cap:%d \n", len(emptyArr), cap(emptyArr))
}
