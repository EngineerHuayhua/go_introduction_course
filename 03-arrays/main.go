package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)
	var myArrayVar [5]int
	fmt.Println(myArrayVar)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	fmt.Println()
	fmt.Println("Array with values")
	myArrayVar1 := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println(myArrayVar1)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	fmt.Println()
	fmt.Println("slice")
	var mySliceVar []int
	fmt.Println(mySliceVar)
	fmt.Printf("Type: %T, bytes: %d, bits: %d\n", mySliceVar, unsafe.Sizeof(mySliceVar), unsafe.Sizeof(mySliceVar)*8)
	mySliceVar = append(mySliceVar, 1, 2, 3, 4, 5)
	fmt.Printf("size: %d, value: %v\n", len(mySliceVar), mySliceVar)

	fmt.Println()
	mySlice := myArrayVar1[2:4] // slice from index 2 to 4 (exclusive)
	fmt.Println(mySlice)
	fmt.Printf("size: %d, value: %v\n", len(mySlice), mySlice)

	fmt.Println(&myArrayVar1[2]) // address of the slice
	fmt.Println(&mySlice[0])     // address of the first element of the slice

	fmt.Println(myArrayVar1)

	mySlice[0] = "three modified"
	mySlice[1] = "four modified"
	fmt.Println(myArrayVar1)

	fmt.Println(myArrayVar1[:4]) // slice from the beginning to index 4 (exclusive)
	fmt.Println(myArrayVar1[1:]) // slice from index 1 to the end

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value: %v\n", len(slice), slice)
}
