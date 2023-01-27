package main

import (
	"fmt"
	"strconv"
	"strings"
)

// insertion sort

func main() {
	var txt string
	fmt.Println("Введите последовательность чисел через запятую: ")
	fmt.Scanln(&txt)
	inputStr := strings.Split(txt, ",")
	numbers := make([]int64, 0, len(inputStr))
	for _, numberStr := range inputStr {
		number, _ := strconv.ParseInt(numberStr, 10, 64)
		numbers = append(numbers, number)
	}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i; j >= 0; j-- {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			} else {
				break
			}
		}
	}
	fmt.Println(numbers)

}
