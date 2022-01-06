//Create MAP perform add, update, delete, and get operation delete all.

package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"shopping": 2000,
		"card":     5000,
		"shops":    20,
	}

	fmt.Println(mp["shopping"]) //Get the items
	mp["shopping"] = 900        //Insert new item
	//delete(mp, "shopping")//Delete the items
	fmt.Println(mp)
}
