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

	ch := make(chan struct{})
	quicksort(numbers, 0, len(numbers)-1, cond1, ch)
	<-ch
	fmt.Println(numbers)

	ch = make(chan struct{})
	quicksort(numbers, 0, len(numbers)-1, cond2, ch)
	<-ch
	fmt.Println(numbers)

}

func cond1(a int, b int) bool {
	return a > b
}

func cond2(a int, b int) bool {
	return a < b
}

func quicksort(arr []int, lo int, hi int, cond func(a int, b int) bool, ch chan struct{}) {
	defer close(ch)
	if lo < hi {
		p := partition(arr, lo, hi, cond)
		ch1 := make(chan struct{})
		go quicksort(arr, lo, p, cond, ch1)
		ch2 := make(chan struct{})
		go quicksort(arr, p+1, hi, cond, ch2)
		<-ch1
		<-ch2
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
