package main

import "fmt"

type Stack struct {
	stack []string
}

func (s *Stack) push(char string) {
	s.stack = append(s.stack, char)
}

func (s *Stack) pop() string {
	lastElement := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return lastElement
}

func (s Stack) read() string {
	return s.stack[len(s.stack)-1]
}

func Reverse(str string) string {
	stack := Stack{make([]string, 0)}

	for _, char := range str {
		stack.push(string(char))
	}

	reverseStr := ""
	stackLength := len(stack.stack)

	for i := 0; i < stackLength; i++ {
		reverseStr += stack.pop()
	}

	return reverseStr
}

func main() {
	// s1 := Stack{make([]string, 0)}
	// fmt.Println(s1.stack)
	// s1.push("a")
	// fmt.Println(s1.stack)
	// s1.push("c")
	// s1.push("e")
	// fmt.Println(s1.stack)
	// fmt.Println(s1.pop())
	// fmt.Println(s1.stack)
	// fmt.Println(s1.read())
	fmt.Println(Reverse("abcde"))
}
