package main

import ("encoding/json"
"fmt"
"io/ioutil"
"os"
// "strconv"
)
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

//linked list

type Node struct {
    Value interface{}
    next  *Node
}

type List struct {
    head *Node
    size int
}

func (list *List) Insert(Value interface{}) {

    node := Node{Value, nil}
    if list.head == nil {
        list.head = &node
    } else {
        temp := list.head

        for temp.next != nil {
            temp = temp.next
        }

        temp.next = &node
    }
    list.size++
}

func (list *List) Remove(Value interface{}) {
    if list.head == nil {
        fmt.Println("List is empty")
    } else {
        temp := list.head

        if temp.Value == Value {
            list.head = temp.next
        } else {
            for {
                if temp.next.Value == Value  {
                    break
                }
                temp = temp.next
            }
            temp.next = temp.next.next
        }
    }
    list.size--
}
// func (list *List) get(value int) int{   
//     node:=list.head
//     var r int =0
//     for i:=0;i<value;i++{
//         if i == value-1{
//             r=list.Value
//         }
//         node=node.next
//     }
//     return r
// }
func (list *List) display() {
    temp := list.head
    for temp.next != nil {
        fmt.Println(temp.Value)
        temp = temp.next
    }
    fmt.Println(temp.Value)
}
func main(){     

	type User struct {
		Name   string `json:"name"`
		Symbol   string `json:"symbol"`
		Share    int    `json:"share"`
	}
	type Users struct{
		Users []User `json:"users"`
	}
    s:=new(Stack)
    link:=new (List)
	//open te json file
    jsonfile,err:=os.Open("filestack.json")
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

    for i := 0; i < len(users.Users); i++ {
        l:=users.Users[i].Symbol
        link.Insert(l)
    }
    for i:=0;i<link.size;i++{

        s.push(i)
    }
    for i:=0;i<link.size;i++{
      fmt.Println(s.pop())     
    }
    
}