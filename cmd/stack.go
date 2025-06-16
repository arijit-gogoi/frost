package main

import "log"

type Stack struct {
	stack []Word
}

// Push adds an item to the top of the stack
func (s *Stack) Push(val Word) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (top Word) {
	top = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Top() Word {
	if len(s.stack) < 0 {
		log.Fatal("len(s.stack) is 0")
	}
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Second() Word {
	return s.stack[len(s.stack)-2]
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Clear() {
	s.stack = []Word{}
}
