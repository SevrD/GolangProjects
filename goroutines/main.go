package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Написать соритровку слайса по произвольному условию (func sort(arr []int, cond func(a int, b int) bool))

func main() {

	n := 10000000

	var numbers []int
	for i := 0; i < n; i++ {
		numbers = append(numbers, int(rand.Intn(100)))
	}
	// fmt.Println("Исходный массив: ")
	// fmt.Println(numbers)
	// fmt.Println("Отсортированный массив: ")
	fmt.Println(time.Now())
	sort.Ints(numbers)
	fmt.Println(time.Now())
	//fmt.Println(numbers)

	numbers = nil
	for i := 0; i < n; i++ {
		numbers = append(numbers, int(rand.Intn(100)))
	}
	fmt.Println(time.Now())
	var wg sync.WaitGroup
	wg.Add(1)
	go quicksort(numbers, 0, len(numbers)-1, cond1, &wg)
	wg.Wait()
	fmt.Println(time.Now())
	//fmt.Println(numbers)

	// numbers = nil
	// for i := 0; i < 200; i++ {
	// 	numbers = append(numbers, int(rand.Intn(100)))
	// }
	// fmt.Println("Исходный массив: ")
	// fmt.Println(numbers)
	// wg.Add(1)
	// go quicksort(numbers, 0, len(numbers)-1, cond2, &wg)
	// wg.Wait()
	// fmt.Println("Отсортированный массив: ")
	// fmt.Println(numbers)
	fmt.Println("Done")
}

func cond1(a int, b int) bool {
	return a > b
}

func cond2(a int, b int) bool {
	return a < b
}

func quicksort(arr []int, lo int, hi int, cond func(a int, b int) bool, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if lo < hi {
		p := partition(arr, lo, hi, cond)
		if hi-lo > 50000 {
			if wg != nil {
				wg.Add(1)
			}
			go quicksort(arr, lo, p, cond, wg)
		} else {
			quicksort(arr, lo, p, cond, nil)
		}
		quicksort(arr, p+1, hi, cond, nil)

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
