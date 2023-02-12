package main

import "fmt"

func main() {
	var sl []chan int

	for i := 0; i < 5; i++ {
		sl = append(sl, make(chan int, 1))
		sl[i] <- i
	}
	rs := sum(sl)
	for k := range rs {
		fmt.Println(k)
	}

}

func sum(chls []chan int) chan int {
	result := make(chan int, 5)
	defer close(result)
	var temp int
	for _, ch := range chls {
		temp = <-ch
		result <- temp
	}
	return result
}
