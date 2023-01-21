package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &number)
	fmt.Println("Последовательность Коллатца рекурсией:")
	collatz(number)

	fmt.Println("Последовательность Коллатца циклом:")
	for number > 1 {
		fmt.Print(number, ", ")
		if number%2 == 0 {
			number = number / 2
		} else {
			number = number*3 + 1
		}

	}
	fmt.Println(number)
}

func collatz(n int) {
	fmt.Print(n)
	if n == 1 {
		fmt.Println()
		return
	}
	fmt.Print(", ")
	if n%2 == 0 {
		collatz(n / 2)
	} else {
		collatz(n*3 + 1)
	}
}
