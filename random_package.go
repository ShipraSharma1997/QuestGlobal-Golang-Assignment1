//Take a slice appended element in the slice using rand packages

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var reason string
	//create the reasons slice and append reasons to it
	reasons := make([]string, 0)
	reasons = append(reasons,
		"Locked out",
		"Pipes broke",
		"Food poisoning",
		"Not feeling well")
	rand.Seed(time.Now().Unix())
	message := fmt.Sprint("Gonna work from home...", reasons[rand.Intn(len(reasons))])
	fmt.Println(reason)
	fmt.Println(message)

}
