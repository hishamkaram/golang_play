package main

import (
	"fmt"
	"time"
)

// RoutineFunc is a routine function
func RoutineFunc(caller string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s \t-\t %d \t on: %s \n", caller, i, time.Now().Format(time.RFC850))
	}
}
func main() {
	go RoutineFunc("caller1")
	go RoutineFunc("caller2")
	fmt.Println("Press <enter> key to close")
	fmt.Scanln()
	fmt.Println("Terminated")
}
