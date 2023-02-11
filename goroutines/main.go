package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(1)
	go quicksort(numbers, 0, len(numbers)-1, cond1, &wg)
	wg.Wait()
	fmt.Println(numbers)
	wg.Add(1)
	go quicksort(numbers, 0, len(numbers)-1, cond2, &wg)
	wg.Wait()
	fmt.Println(numbers)

}

func cond1(a int, b int) bool {
	return a > b
}

func cond2(a int, b int) bool {
	return a < b
}

func quicksort(arr []int, lo int, hi int, cond func(a int, b int) bool, wg *sync.WaitGroup) {
	defer wg.Done()
	if lo < hi {
		p := partition(arr, lo, hi, cond, wg)
		go quicksort(arr, lo, p, cond, wg)
		go quicksort(arr, p+1, hi, cond, wg)
	}
}

func partition(arr []int, low int, hight int, cond func(a int, b int) bool, wg *sync.WaitGroup) int {
	wg.Add(2)
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
