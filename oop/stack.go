package main

import "fmt"

type Stack struct {
	list []int
}

func (s *Stack) pop() int {
	i := len(s.list)
	elm := s.list[i-1]
	s.list = s.list[:i-1]
	return elm
}

func (s *Stack) push(val int) {
	s.list = append(s.list, val)
}

func main() {
	st := Stack{[]int{1, 2, 3}}
	//fmt.Println(st.pop())
	st.push(5)
	fmt.Println(st.list)
}
