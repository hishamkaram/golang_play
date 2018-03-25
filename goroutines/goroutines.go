package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var c1 = make(chan string)

// RoutineFunc is a routine function
func RoutineFunc(caller string) {
	for i := 0; i < 10; i++ {
		msg := []string{caller, strconv.Itoa(i), "on: ", string(time.Now().Format(time.RFC850)), "\n"}
		c1 <- strings.Join(msg, "\t")
	}
}
func main() {
	go RoutineFunc("caller1")
	go RoutineFunc("caller2")
	for i := 0; i < 20; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("received - \t %s", msg1)
		}
	}

	fmt.Println("Press <enter> key to close")
	fmt.Scanln()
	fmt.Println("Terminated")
}
