package main

import (
	"fmt"
	"strings"
)

var strArr = []string{"Hello", "World!"}

func main() {
	fmt.Println(strings.Join(strArr, " "))
	for _, chr := range []rune(strings.Join(strArr, " ")) {
		fmt.Printf("'%d' \t-\t %c \n", chr, chr)
	}
}
