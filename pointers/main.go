package main

import "fmt"

func main() {
	// var ptr *int
	// type of pointer must match type of variable

	var ptr *int
	var a int = 10
	ptr = &a

	fmt.Println(a)
	fmt.Println(*ptr) // dereferencing a pointer

	modifyValue(ptr)

	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
