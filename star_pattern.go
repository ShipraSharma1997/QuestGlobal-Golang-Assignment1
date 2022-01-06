// Start Patter Program (Inverted half pyramid of *)

package main

import "fmt"

func main() {
	var rows int
	fmt.Printf("Enter the number of rows: ")
	fmt.Scanf("%d", &rows)
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
	return
}
