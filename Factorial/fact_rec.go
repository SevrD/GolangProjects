package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return n
	}
	return n * fact(n-1)
}

func main() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Factorial: %d", fact(number))
}
