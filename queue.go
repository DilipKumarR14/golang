package main
import("fmt"
"strings")
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
func main(){

	var q *Queue=new (Queue)
	var str string
	fmt.Println("Enter the expression : ")
	fmt.Scanf("%s",&str)
	res:=strings.Split(str,"")
	// fmt.Println(res)
	for i:=0;i<len(res);i++{
		q.enqueue(res[i])
	}
	fmt.Println(q)
	var str1 string=""

	for i:=len(res)-1;i>=0;i--{
		str1 = q.dequeue() + str1
	}
	fmt.Println(str1)
	if str == str1{
		fmt.Println("palindrome\n")
	}else{
		fmt.Println("not a palindrome\n")
	}


}