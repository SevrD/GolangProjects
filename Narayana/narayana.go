package main

import (
	"fmt"
	"sort"
	"strings"
)

// Написать функцию, которая по слайсу будет выводить все возможные перестановки его элементов.
// Для 1,2,3 это будет 1,2,3 1,3,2 2,1,3 2,3,1 3,1,2 3,2,1
func main() {

	var txt string
	fmt.Println("Введите последовательность чисел через запятую: ")
	fmt.Scanln(&txt)
	input := strings.Split(txt, ",")
	sort.Strings(input)
	fmt.Println("Возможные комбинации: ")
	fmt.Println(strings.Join(input, ","))
	for i, j := len(input)-2, len(input)-1; i >= 0; i, j = maxI(input) {
		input[i], input[j] = input[j], input[i]
		for k, l := i+1, len(input)-1; k < l; k++ {
			input[k], input[l] = input[l], input[k]
			l--
		}
		fmt.Println(strings.Join(input, ","))
	}

}

func maxI(sequence []string) (int, int) {
	for i := len(sequence) - 2; i >= 0; i-- {
		if sequence[i] < sequence[i+1] {
			for j := len(sequence) - 1; j > i; j-- {
				if sequence[i] < sequence[j] {
					return i, j
				}
			}
		}
	}
	return -1, -1
}
