package main

import "fmt"

func main() {

	// sequence := adder() // i initialize to 0 first
	// //  then i won't reset
	// fmt.Println(sequence()) // i = 1
	// fmt.Println(sequence()) // i = 2
	// fmt.Println(sequence()) // i = 3

	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtractor := func() func(int) int {
		countDown := 99
		return func(i int) int {
			countDown -= i
			return countDown
		}
	}()

	// using the closure subtractor
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))

}

func adder() func() int { // func() int - this is return function
	i := 0 // i use inmemory
	fmt.Println("previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}

}
