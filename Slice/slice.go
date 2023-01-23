package main

import (
	"fmt"
	"strings"
)

// Написать функцию, которая по слайсу будет выводить все возможные перестановки его элементов.
// Для 1,2,3 это будет 1,2,3 1,3,2 2,1,3 2,3,1 3,1,2 3,2,1
func main() {

	var txt string
	fmt.Println("Введите последовательность чисел через запятую: ")
	fmt.Scanln(&txt)
	input := strings.Split(txt, ",")
	var result [][]string = make([][]string, 0)

	fill(&result, input, make([]string, 0))

	fmt.Println("Возможные комбинации: ")
	for _, value := range result {
		fmt.Println(strings.Join(value, ","))
	}

}

func fill(result *[][]string, in []string, out []string) {

	for i, number := range in {
		newOut := append(out, number)
		slice := make([]string, 0, len(in)-1)
		slice = append(slice, in[:i]...)
		slice = append(slice, in[i+1:]...)
		fill(result, slice, newOut)
	}
	if len(in) == 0 {
		*result = append(*result, out)
	}
}
