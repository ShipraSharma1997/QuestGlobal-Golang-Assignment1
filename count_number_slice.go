// Find out How many times each number accesses in slices using map.

// Find out How many times each number accesses in slices using map.

package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"

	v := make([]string, 0, len(m))

	for _, value := range m {
		v = append(v, value)
	}
	fmt.Println(len(m))
}
