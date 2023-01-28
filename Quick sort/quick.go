package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Написать соритровку слайса по произвольному условию (func sort(arr []int, cond func(a int, b int) bool))

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
	fmt.Println(sort(numbers, cond1))
	fmt.Println(sort(numbers, cond2))
}

func cond1(a int64, b int64) bool {
	return a > b
}

func cond2(a int64, b int64) bool {
	return a < b
}

func sort(arr []int64, cond func(a int64, b int64) bool) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	base := len(arr) / 2
	left := make([]int64, 0)
	right := make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		if i == base {
			continue
		}
		if cond(arr[i], arr[base]) {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	if len(left) > 1 {
		left = sort(left, cond)
	}
	if len(right) > 1 {
		right = sort(right, cond)
	}
	left = append(left, arr[base])
	return append(left, right...)
}
