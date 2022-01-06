// Biggest number by using slice

package main

import "fmt"

func main() {

	var largerNumber, temp int
	size := 0
	fmt.Print("Number of elements n=")
	fmt.Scanln(&size)
	fmt.Println("Enter of array elements")
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&elements[i])
	}
	fmt.Println("Entered Array of elements:", elements)
	for _, element := range elements {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}
	fmt.Println("Largest number of Array/slice is  ", largerNumber)
}
