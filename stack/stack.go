package stack


type (
	Stack[T any] struct {
		length int
		head *node[T]
	}
	node[T any] struct {
		value T
		prev *node[T]
	}
)

func NewStack[T any]() *Stack[T] {
	return &Stack[T] {0, nil}
}

func (s *Stack[T]) Push(value T){
	prev := s.head
	newNode := node[T]{value, prev}
	s.head = &newNode
	s.length += 1
}

func (s *Stack[T]) Pop() *node[T] {
	if s.length == 0 {
		return nil
	}
	tmp := s.head
	s.head = tmp.prev
	s.length -= 1
	return tmp
}

func (s *Stack[T]) Length() int {
	return s.length
}
