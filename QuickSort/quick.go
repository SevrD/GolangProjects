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
	sort2(numbers, cond1)
	fmt.Println(numbers)

	sort2(numbers, cond2)
	fmt.Println(numbers)
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
	left := make([]int64, 0, len(arr)-1)
	right := make([]int64, 0, len(arr)-1)
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

func sort2(arr []int64, cond func(a int64, b int64) bool) {
	if len(arr) <= 1 {
		return
	}
	base := len(arr) / 2
	var l, r, lf, rf int
	r = len(arr) - 1

	for l < r {
		lf, rf = -1, -1
		for ; l < base; l++ {
			if cond(arr[base], arr[l]) {
				lf = l
				break
			}
		}
		for ; r > base; r-- {
			if cond(arr[r], arr[base]) {
				rf = r
				break
			}
		}
		if lf >= 0 && rf > 1 {
			arr[lf], arr[rf] = arr[rf], arr[lf]
		} else if lf >= 0 {
			for i := lf; i < base; i++ {
				if cond(arr[base], arr[i]) {
					move(arr, i, len(arr)-1)
					base--
					i--
				}
			}
			break
		} else if rf >= 0 {
			for i := rf; i > base; i-- {
				if cond(arr[i], arr[base]) {
					move(arr, i, 0)
					base++
					i++
				}
			}
			break
		} else {
			break
		}

	}
	if base > 0 {
		sort2(arr[:base], cond)
	}
	if base < len(arr)-1 {
		sort2(arr[base+1:], cond)
	}
}

func move(arr []int64, from int, to int) {
	if from < to {
		value := arr[from]
		for i := from; i < to; i++ {
			arr[i] = arr[i+1]
		}
		arr[to] = value
	} else if from > to {
		value := arr[from]
		for i := from; i > to; i-- {
			arr[i] = arr[i-1]
		}
		arr[to] = value
	}
}
