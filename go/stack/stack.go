package main

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack
	stack.Push(10)
	stack.Push(20)
	if val, ok := stack.Pop(); ok {
		fmt.Println(val) // 20
	}
	if val, ok := stack.Pop(); ok {
		fmt.Println(val) // 10
	}
	if _, ok := stack.Pop(); !ok {
		fmt.Println("Stack isEmpty!")
	}
}
