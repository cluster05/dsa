package stack

type Stack struct {
	items  []int
	length int
}

var (
	InvalidIndex = -999
)

func NewStack() *Stack {
	return &Stack{items: []int{}, length: 0}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	s.length = s.length + 1
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return InvalidIndex, false
	}
	top := s.items[s.length-1]

	s.items = s.items[:s.length-1]
	s.length = s.length - 1
	return top, true
}

func (s *Stack) Top() (int, bool) {
	if s.IsEmpty() {
		return InvalidIndex, false
	}
	return s.items[s.length-1], true
}

func (s *Stack) ElementAt(pos int) (int, bool) {
	if s.IsEmpty() {
		return InvalidIndex, false
	}
	return s.items[s.length-pos], true
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}
