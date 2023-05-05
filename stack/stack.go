package stack

import "errors"

type Stack struct {
	Items []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) <= 0
}

func (s *Stack) Push(item string) {
	if s.IsEmpty() {
		s.Items = append(s.Items, item)
	} else {
		var newItems []string
		newItems = append(newItems, item)
		s.Items = append(newItems, s.Items...)
	}
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("item is empty")
	}

	item := s.Items[0]
	s.Items = s.Items[1:]
	return item, nil
}

func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("item is empty")
	}

	return s.Items[0], nil
}
