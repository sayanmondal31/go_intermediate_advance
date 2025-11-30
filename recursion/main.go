package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println("Sum of digits->", sumOfDigits(2354))

}

// factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// fmt.Println("n->", n)xs

	if n < 10 {
		return n
	}

	// fmt.Println("n%10->", n%10, "| n/10->", n/10)
	return n%10 + sumOfDigits(n/10)
}
