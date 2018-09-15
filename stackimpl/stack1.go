package stackimpl
//stack
type Stack struct{
	top *Element
	size int
}
type Element struct{
	value interface{}
	next *Element
}
// find the size of stack
func (s *Stack) Len() int{
	return s.size
}
// pushing the elements
func (s *Stack) push(value interface{}){
	s.top=&Element{value,s.top}
	s.size++
}
func (s *Stack) pop() (value interface{}){
	if s.size > 0{
		value,s.top=s.top.value,s.top.next
		s.size--
		return
	}
	return nil
}


