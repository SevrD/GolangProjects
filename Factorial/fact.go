package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("Factorial recursion: %d\n", fact(number))

	result := 1
	for i := 2; i <= number; i++ {
		result *= i
	}
	fmt.Printf("Factorial loop: %d\n", result)
}
