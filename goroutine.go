//create 2 function, take global variable i and 2 goroutine should increment value of i untill recahes 100

package main

import (
	"fmt"
	"time"
)

func display(a int) {
	for w := 0; w < 100; w++ {
		time.Sleep(1 * time.Second)
		a = a + 1
		fmt.Println(a)
	}
}

func main() {
	i := 0
	// Calling Goroutine
	go display(i)

	// Calling normal function
	display(i)
}
