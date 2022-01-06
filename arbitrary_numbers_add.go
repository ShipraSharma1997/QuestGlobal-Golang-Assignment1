package main

import (
	"fmt"
	"unsafe"
)

func main() {
	num1 := int64(12345)
	num2 := int64(12345)
	fmt.Print(num1, " :num1 ")
	fmt.Print(num2, " :num2")
	additionArr := make([]byte, 5)
	byteArr1 := IntToByteArray(num1)
	byteArr2 := IntToByteArray(num2)
	fmt.Print(byteArr1, " :byteArr1 ")
	fmt.Print(byteArr2, " :byteArr2")
	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Print(byteArr1[i], " :arr1 ")
		fmt.Print(byteArr2[i], " :arr2 ")
		fmt.Print(additionArr[i], "  ")
		additionArr[i] = byteArr1[i] + byteArr2[j]
		fmt.Print(additionArr[i], "  ")
	}

	numAgain := ByteArrayToInt(additionArr)
	fmt.Println("number:", numAgain)

}

func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
