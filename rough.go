package main
import ("encoding/json"
"fmt"
"io/ioutil"
"os")
// "strconv"
// "log")
// "bytes")
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

	type Users struct{
		Users []User `json:"users"`
	}
	type User struct {
		Name   string `json:"name"`
		Symbol   string `json:"symbol"`
		Share    int    `json:"share"`
	}
// 	var num int
// 	var name string
// 	var sym string
// 	var share int

// 	fmt.Println("enter the no of stocks ")
// 	fmt.Scanf("%d",&num)
// 	for i:=0;i<num;i++{
// 		fmt.Println("enter the name")
// 		fmt.Scanf("%s",&name)

// 		fmt.Println("enter the symbol")
// 		fmt.Scanf("%s",&sym)
		
// 		fmt.Println("enter the noofshare")
// 		fmt.Scanf("%d",&share)

// 		message := User { Name:name,Symbol:sym,Share:share}
// 	buf:=new (bytes.Buffer)
// 	encoder:=json.NewEncoder(buf)
// 	encoder.Encode(message)
// 	file, err :=os.Create("res.json")
// 	checkerror(err)
// 	defer file.Close()
// 	io.Copy(file,buf)
		
// 	}
	
// }
// 	func checkerror(err error){
// 		if err!=nil{
// 			log.Fatal(err)
// 		}
// 	}	
	

	 //open te json file
	 jsonfile,err:=os.Open("file.json")
	 // if we os.open returns error
	 if err!=nil{
		 fmt.Println(err)
	 }
	 fmt.Println("file opened")
	 defer jsonfile.Close()
	 //read a file
	 bytevalue,_:=ioutil.ReadAll(jsonfile)
 
	 // initialize user array
 
	 var users Users
 
	 // we unmarshall our byte array which contains our jsonfile
 
	 json.Unmarshal(bytevalue,&users)
 
	s:=new(Stack)

	for i := 0; i < len(users.Users); i++ {
        fmt.Println("StockName : " + users.Users[i].Name)
		fmt.Println("Symbol : " + users.Users[i].Symbol)
        fmt.Println("Share : " + strconv.Itoa(users.Users[i].Share))// integer format int64
        s.push(users.Users[i].Symbol)
        fmt.Println()
    }

    
}