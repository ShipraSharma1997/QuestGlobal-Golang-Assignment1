//Write a multidimenstional array, and assign values and try to access values an try with rang and without range.
package main

import "fmt"

//With range Keyword
func main() {

	input := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println("Using range Keyword")
	for _, element := range input {
		for _, val := range element {
			fmt.Println(val)
		}
	}
}

//Without range Keyword
/*func main() {

    input := [2][3]int{{1,2,3},{2,3,4}}
       fmt.Println("Without Using range Keyword")
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Println(input[i][j])
        }
    }
}*/
