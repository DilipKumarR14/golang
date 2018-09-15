package main
import "fmt"
import "github.com/utility"
//stack
type Stack struct{
	top *Element
	size int
}
//LinkedList
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
func reverse(n int) bool {
    new_int := 0
    for n > 0 {
        remainder := n % 10
        new_int *= 10
        new_int += remainder 
        n /= 10
    }
	if new_int == n {
		return true
	} else{
		return false
	}
}
func prime(n1 int ) bool {
	for i:=2; i<=100; i++{
        isPrime:=true
        for j:=2; j<i; j++{
            if i % j == 0 {
                return false
            }
        }
        if isPrime == true{
			return true
		}
	}
	return true
}

func main(){
	
	s:=new(Stack)
	for i:=1;i<100;i++{
		if(prime(i) && reverse(i)){
			s.push(i)
		}
	}
	for s.Len()>0 {
		fmt.Println(s.pop())
	}
utility.Add()
}   