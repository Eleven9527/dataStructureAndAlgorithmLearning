package stack

import "fmt"

/*
	栈demo
*/

type stack struct {
	Top  int //栈顶元素的索引
	List []int
}

//初始化栈
func initStack() *stack {
	return &stack{
		Top:  -1,
		List: make([]int, 0),
	}
}

//入栈
func (s *stack) push(val int) {
	s.Top++
	s.List = append(s.List, val)
}

//出栈,bool=false表示栈已空
func (s *stack) pop() (int, bool) {
	if s.Top == -1 {
		return 0, false
	}

	v := s.List[s.Top]
	s.Top--

	return v, true
}

func StackTest(vals []int) {
	fmt.Println("vals = ", vals)
	//init
	s := initStack()

	//push
	for _, v := range vals {
		s.push(v)
	}

	//pop
	for {
		if v, ok := s.pop(); ok {
			fmt.Println("出栈元素 = ", v)
		} else {
			break
		}
	}
}
