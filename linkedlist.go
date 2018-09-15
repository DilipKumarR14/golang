package main
 
import ("fmt"
"io/ioutil"
"strings"
"os"
)

type Node struct {
    Value string
    next  *Node
}

type List struct {
    head *Node
    size int
}

func (list *List) Insert(Value string) {

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
func (list *List) searching(search string) string {
    temp:=list.next
    if temp.head == nil{
        fmt.Println("empty list")
    }else{
        
        for temp!=nil{
            if temp.Value == search{
                return temp.Value
            }else{
                temp=temp.next
            }
        }
    }
    return temp.Value
}
func (list *List) display() {
    temp := list.head
    for temp.next != nil {
        fmt.Println(temp.Value)
        temp = temp.next
    }
    fmt.Println(temp.Value)
}

func main() {
    link := List{}
    b, err := ioutil.ReadFile("file.txt")
    if err != nil {
        fmt.Print(err)
    }
    str := string(b) // convert content to a 'string'
    character :=strings.Fields(str)
    for _,v:=range character{
        link.Insert(v)
    }

    link.display()  

    fmt.Print("ENter the strings to search : ")
    var searching string
    fmt.Scanf("%s",&searching)
    fmt.Println()
    for i:=0;i<link.size;i++{
        if link.searching(searching) != searching{
            fmt.Println("not found added")
            f,err:=os.OpenFile("file.txt",os.O_APPEND|os.O_WRONLY,0600)// open the file for writing appending
            if err!=nil{
                panic(err)
            }
            defer f.Close()
            if _,err = f.WriteString(searching);err!=nil{
                panic(err)
            }
            break  
        }else{
            fmt.Println("found deleted")
            f,err:=os.OpenFile("file.txt",os.O_APPEND|os.O_WRONLY,0600)// open the file for writing appending
            if err!=nil{
                panic(err)
            }
            defer f.Close()
            link.Remove(searching)

        }
    }

    



//   link.Insert(10)
//   link.Insert(20)
}