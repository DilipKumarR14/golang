package l 
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

// func main() {
//     link := List{}
    
//     }

    



//   link.Insert(10)
//   link.Insert(20)
}