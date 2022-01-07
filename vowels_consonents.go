//Take a string using switch findout number of vovels, consonenets and special char

package main

import "fmt"

func main() {  
 str := "kiran@"  
 vowel := 0
 consonent := 0
 special_character := 0
 
 for _, ch := range str {  
  switch ch {  
  case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':  
   vowel++  
   
   case '!', '@', '#', '$', '%', '^', '&', '*','(', ')':  
   special_character++
   
    case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k','l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'w', 'v','x','y','z':  
   consonent++
 
  }
 
 }  
 fmt.Printf(" %s string contains vowels count: %d\n", str, vowel)  
 fmt.Printf(" %s string contains constant count: %d\n", str, consonent)
    fmt.Printf(" %s string contains special character count: %d\n", str, special_character)  


 
}  
