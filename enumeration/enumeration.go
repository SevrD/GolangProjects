// Перебором найти два целых числа между которыми находиться корень заданного
// (например для 15 это 3 (9 меньше 15) и 4 (а 16 уже больше)) (попробовать и рекурсией и циклом)

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите целое число: ")
	fmt.Scanf("%d", &number)
	second := 1
	for {
		if second*second > number {
			break
		}
		second += 1
	}
	fmt.Printf("Корень заданного числа находится между %d и %d (циклом)\n", second-1, second)
	second = find(1, number)
	fmt.Printf("Корень заданного числа находится между %d и %d (рекурсией)", second-1, second)
}

func find(index int, number int) int {
	if index*index > number {
		return index
	}
	return find(index+1, number)
}
