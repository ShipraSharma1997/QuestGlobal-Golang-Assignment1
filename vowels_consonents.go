//Take a string using switch findout number of vovels, consonenets and special char

package main

import "fmt"

func main() {
	str := "kiran@"
	count := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':

			count++
		}

	}
	fmt.Printf(" %s string contains vowels count: %d\n", str, count)
	fmt.Printf(" %s string contains constant count: %d\n", str, count)
	fmt.Printf(" %s string contains special character count: %d\n", str, count)

}
