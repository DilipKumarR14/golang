package main

import (
    "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

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

func (list *List) display() {
    temp := list.head
    for temp.next != nil {
        fmt.Println(temp.Value)
        temp = temp.next
    }
    fmt.Println(temp.Value)
}
// user struct which contains array of users
type Users struct{
    Users []User `json:"users"`
}
type User struct {
	Name   string `json:"name"`
	Symbol   string `json:"symbol"`
	Share    int    `json:"share"`
}
func main() {

    list := new(List)
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

    for i := 0; i < len(users.Users); i++ {
        fmt.Println("StockName : " + users.Users[i].Name)
		fmt.Println("Symbol : " + users.Users[i].Symbol)
        fmt.Println("Share : " + strconv.Itoa(users.Users[i].Share))
        list.Insert(users.Users[i])
        fmt.Println()
    }
        list.display()
        fmt.Println("enter the stock to be deleted")
        fmt.Println("0-airtel / 1-jio / 2-idea")
        var option int
        fmt.Scanf("%d",&option)
        list.Remove(users.Users[option])
        
       
        fmt.Println()
        list.display()

          //marshalling 

        // stma:=User{}
        // stmajson, err :=json.Marshal(stma)
        // if err!=nil{
        //     fmt.Println(err)
        // }
        // err=ioutil.WriteFile("file.json",stmajson,0644)
        // fmt.Printf("%+v\n", stma)
        // fmt.Println(stmajson)


}

