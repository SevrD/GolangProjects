package main

import (
	"fmt"
	"math/rand"
)

// Написать соритровку слайса по произвольному условию (func sort(arr []int, cond func(a int, b int) bool))

func main() {

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, int(rand.Intn(100)))
	}
	fmt.Println("Исходный массив: ")
	fmt.Println(numbers)
	fmt.Println("Отсортированный массив: ")

	quicksort(numbers, 0, len(numbers)-1, cond1)
	fmt.Println(numbers)

	quicksort(numbers, 0, len(numbers)-1, cond2)
	fmt.Println(numbers)

}

func cond1(a int, b int) bool {
	return a > b
}

func cond2(a int, b int) bool {
	return a < b
}

func quicksort(arr []int, lo int, hi int, cond func(a int, b int) bool) {
	if lo < hi {
		p := partition(arr, lo, hi, cond)
		quicksort(arr, lo, p, cond)
		quicksort(arr, p+1, hi, cond)
	}
}

func partition(arr []int, low int, hight int, cond func(a int, b int) bool) int {
	pivot := arr[(hight+low)/2]
	for {
		for cond(arr[low], pivot) {
			low++
		}
		for cond(pivot, arr[hight]) {
			hight--
		}
		if low >= hight {
			return hight
		}
		arr[low], arr[hight] = arr[hight], arr[low]
		low++
		hight--
	}
}
