package valid_parentheses

import (
	"errors"
)

type stack struct {
	items []string
}

func (s *stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() (string, error) {
	len := len(s.items)

	if len == 0 {
		return "", errors.New("no data")
	}
	item, items := s.items[len-1], s.items[0:len-1]
	s.items = items

	return item, nil
}

func isValid(s string) bool {
	runes := []rune(s)
	stack := stack{items: []string{}}

	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case '(':
			stack.Push("(")
		case '{':
			stack.Push("{")
		case '[':
			stack.Push("[")
		case ')':
			if len(stack.items) == 0 || stack.items[len(stack.items)-1] != "(" {
				return false
			}
			s, err := stack.Pop()
			if err != nil {
				return false
			}

			if s != "(" {
				return false
			}
			continue
		case '}':
			if len(stack.items) == 0 || stack.items[len(stack.items)-1] != "{" {
				return false
			}
			s, err := stack.Pop()
			if err != nil {
				return false
			}

			if s != "{" {
				return false
			}
			continue
		case ']':
			if len(stack.items) == 0 || stack.items[len(stack.items)-1] != "[" {
				return false
			}
			s, err := stack.Pop()
			if err != nil {
				return false
			}

			if s != "[" {
				return false
			}
			continue
		}

	}
	if len(stack.items) != 0 {
		return false
	}
	return true
}
