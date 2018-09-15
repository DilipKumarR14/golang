package main
import ("encoding/json"
// "fmt"
"io"
"os"
// "strconv"
"log"
"bytes")
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
// user struct which contains array of users

func main(){

	// type Users struct{
	// 	Users []User `json:"users`
	// }
	type User struct {
		Name   string `json:"name"`
		Symbol   string `json:"symbol"`
		Share    int    `json:"share"`
	}
	message := User { Name:"airtel",Symbol:"!",Share:2}
	buf:=new (bytes.Buffer)
	encoder:=json.NewEncoder(buf)
	encoder.Encode(message)
	file, err :=os.Create("res.json")
	checkerror(err)
	defer file.Close()
	io.Copy(file,buf)
}
	func checkerror(err error){
		if err!=nil{
			log.Fatal(err)
		}
	}	
	
// 	s:=new(Stack)

// 	for i := 0; i < len(users.Users); i++ {
//         fmt.Println("StockName : " + users.Users[i].Name)
// 		fmt.Println("Symbol : " + users.Users[i].Symbol)
//         fmt.Println("Share : " + strconv.Itoa(users.Users[i].Share))// integer format int64
//         s.push(users.Users[i].Symbol)
//         fmt.Println()
//     }

//     for s.Len()>0 {
// 		fmt.Println(s.pop())
//     }
// }