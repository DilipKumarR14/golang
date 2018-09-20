package main
import ("fmt"
"math/rand"
"time"
"utility"
)

func main(){
	link := List{}
	cards()
}

func cards(){
	suit:=[4]string{"Clubs","Diamonds", "Hearts", "Spades"}
	rank:=[13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10","Jack", "Queen", "King", "Ace"}
	deck:=[52]string{}
	l:=0
	for i:=0;i<len(suit);i++{
		for j:=0;j<len(rank);j++{
			deck[l]=suit[i]+" of "+rank[j]
			l++
		}
	}
	arr:=[4][9]string{}
	a:=deck
	
	for i:=0;i<len(deck);i++{
		rand.Seed(time.Now().UTC().UnixNano()) 
		r:=randInt(0, 51)
		var temp string
		temp=a[r]
		a[r]=a[i]
		a[i]=temp
	}
	num:=0
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			arr[i][j]=a[num]
			num++	
		}
	}
	fmt.Println(arr)
}
func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}