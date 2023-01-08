package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readInput() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	command := strings.TrimSpace(s)
	if command == "" {
		readInput()
		return
	}
	switch {
	case strings.HasPrefix(command, "/"):
		{
			switch command {
			case "/exit":
				fmt.Println("Bye!")
				return
			case "/help":
				fmt.Println("The program calculates the sum of numbers")
			default:
				fmt.Println("Unknown command")
			}
		}
	case strings.Contains(command, "="):
		readVariable(command)
	case IsLetter(command):
		printVariable(command)
	default:
		if !fillStack(s) {
			fmt.Println("Invalid expression")
		} else {
			printResult()
		}
	}
	readInput()
}

func printVariable(name string) {
	elem, ok := variables[name]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("Unknown variable")
	}
}

func readVariable(expression string) {
	items := strings.Split(expression, "=")
	if len(items) != 2 {
		fmt.Println("Invalid assignment")
		return
	}
	left := strings.TrimSpace(items[0])
	right := strings.TrimSpace(items[1])
	if !IsLetter(left) || (!IsLetter(right) && !IsNumber(right)) {
		fmt.Println("Invalid identifier")
		return
	}

	value, err := strconv.Atoi(right)
	if err != nil {
		elem, ok := variables[right]
		if ok {
			value = elem
		} else {
			fmt.Println("Unknown variable")
			return
		}
	}
	variables[left] = value
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

var variables = make(map[string]int)

type dataStack []itemType

type itemType struct {
	Operation Operation
	Number    float64
}

type Operation int

const (
	Number Operation = iota + 1
	Add
	Subtract
	Multiply
	Division
	Power
	LeftParenthesis
	RightParenthesis
)

type Stack struct {
	cursor int
	data   dataStack
}

var stack Stack
var postFix Stack

func addToStack(stack *Stack, value itemType) {
	(*stack).cursor += 1
	if len((*stack).data) > (*stack).cursor {
		(*stack).data[(*stack).cursor] = value
	} else {
		(*stack).data = append((*stack).data, value)
	}
}

func addOperator(value itemType) {
	if stack.cursor == -1 || stack.data[stack.cursor].Operation == LeftParenthesis ||
		IsHigherPrecedence(value.Operation, stack.data[stack.cursor].Operation) {
		addToStack(&stack, value)
	} else {
		for stack.cursor >= 0 {
			operation := stack.data[stack.cursor].Operation
			if IsHigherPrecedence(value.Operation, operation) || operation == LeftParenthesis {
				break
			}
			item := pop(&stack)
			addToStack(&postFix, item)
		}
		addToStack(&stack, value)
	}
}

func addParenthesis(value itemType) bool {
	if value.Operation == LeftParenthesis {
		addToStack(&stack, value)
	} else {
		for stack.cursor >= 0 {
			if stack.data[stack.cursor].Operation == LeftParenthesis {
				break
			}
			item := pop(&stack)
			addToStack(&postFix, item)
		}
		if stack.cursor < 0 {
			return false
		}
		pop(&stack)
	}
	return true
}

func closeStack() bool {
	for stack.cursor >= 0 {
		item := pop(&stack)
		if item.Operation == LeftParenthesis {
			return false
		}
		addToStack(&postFix, item)
	}
	return true
}

func IsHigherPrecedence(first Operation, second Operation) bool {
	return first == Power && second != Power ||
		(first == Multiply || first == Division) && (second == Subtract || second == Add)
}

func pop(stack *Stack) itemType {
	result := (*stack).data[(*stack).cursor]
	(*stack).data = (*stack).data[:(*stack).cursor]
	(*stack).cursor -= 1
	return result
}

func fillStack(value string) bool {
	var element itemType
	var prevElement itemType
	stack.cursor = -1
	postFix.cursor = -1
	stack.data = nil
	postFix.data = nil
	for value != "" {
		value = strings.TrimSpace(value)
		if value == "" {
			break
		}
		result := value[:1]
		i := 2
		if IsNumber(result) {
			for i <= len(value) && IsNumber(value[:i]) {
				result = value[:i]
				i += 1
			}
			value = value[i-1:]
			number, err := strconv.Atoi(result)
			if err != nil {
				log.Fatal(err)
			}
			element = itemType{
				Operation: Number,
				Number:    float64(number)}
			addToStack(&postFix, element)
		} else if IsLetter(result) {
			for i < len(value) && IsLetter(value[:i]) {
				result = value[:i]
				i += 1
			}
			value = value[i-1:]

			element = itemType{
				Operation: Number,
				Number:    float64(variables[result])}
			addToStack(&postFix, element)
		} else if result == "+" {
			for value[:1] == "+" {
				value = value[1:]
			}
			element = itemType{
				Operation: Add}
			addOperator(element)
		} else if result == "-" {
			count := 0
			for value[:1] == "-" {
				value = value[1:]
				count += 1
			}
			if count%2 == 0 {
				element = itemType{
					Operation: Add}
			} else {
				element = itemType{
					Operation: Subtract}
			}
			addOperator(element)
		} else if result == "*" {
			value = value[1:]
			element = itemType{
				Operation: Multiply}
			addOperator(element)
		} else if result == "/" {
			value = value[1:]
			element = itemType{
				Operation: Division}
			addOperator(element)
		} else if result == "^" {
			value = value[1:]
			element = itemType{
				Operation: Power}
			addOperator(element)
		} else if result == "(" {
			value = value[1:]
			element = itemType{
				Operation: LeftParenthesis}
			addParenthesis(element)
		} else if result == ")" {
			value = value[1:]
			element = itemType{
				Operation: RightParenthesis}
			if !addParenthesis(element) {
				return false
			}
		}
		if prevElement.Operation != 0 && IsOperator(prevElement.Operation) && IsOperator(element.Operation) {
			return false
		}
		prevElement = element
	}
	return closeStack()
}

func IsOperator(operation Operation) bool {
	return operation == Add || operation == Multiply || operation == Subtract || operation == Division ||
		operation == Power
}

func printResult() {
	var result Stack
	result.cursor = -1
	result.data = nil
	var item2 itemType
	for _, item := range postFix.data {
		if item.Operation == Number {
			addToStack(&result, item)
		} else {
			item1 := pop(&result)
			if result.cursor >= 0 {
				item2 = pop(&result)
			} else {
				item2 = itemType{
					Operation: Number,
					Number:    0}
			}

			newItem := itemType{
				Operation: Number,
				Number:    perform(item2.Number, item1.Number, item.Operation),
			}
			addToStack(&result, newItem)
		}
	}
	fmt.Println(result.data[0].Number)
}

func perform(num1 float64, num2 float64, operation Operation) float64 {
	switch operation {
	case Add:
		return num1 + num2
	case Subtract:
		return num1 - num2
	case Multiply:
		return num1 * num2
	case Division:
		return num1 / num2
	case Power:
		return math.Pow(num1, num2)
	}
	return 0
}

func main() {
	readInput()
}
