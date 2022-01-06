//Write a program to reverse a string do not use a libruary function

package main

import "fmt"

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	//Reverse the string
	str := "Shipra"
	//return the reversed string
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
