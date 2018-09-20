package q
import("fmt"
// "strings"
)
type Queue struct{
	slice[] string
	count int
} 
func (q *Queue) enqueue (i string){
	q.slice = append(q.slice, i)
	q.count++
}
func (q *Queue) dequeue() string{
	ret:=q.slice[0]
	q.count--
	q.slice=q.slice[1:len(q.slice)]
	return ret
}
func (q *Queue) String() string{
	return fmt.Sprint(q.slice)
}
// func main(){

// 	var q *Queue=new (Queue)
// 	q.enqueue("sad")
// 	q.enqueue("das")
// 	q.enqueue("asd")
// 	fmt.Println(q)
// }