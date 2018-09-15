package main
import ("fmt"
// "math/rand"
)

func main(){
	
	cards()
	// deck()
}
func cards() []string{
	suit:=[4]string{"Clubs","Diamonds", "Hearts", "Spades"}
	rank:=[13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10","Jack", "Queen", "King", "Ace"}
	deck:=[52]string{}
	l:=0
	for i:=0;i<len(suit);i++{
		for j:=0;j<len(rank);j++{
			deck[l]=suit[i]+" of "+rank[j]
			
			fmt.Println(deck[l])
			fmt.Println(l)
			l++
		}
	}
	fmt.Println(deck)
	return []deck
}
// func random(min int, max int) int {
//     return rand.Intn(max-min) + min
// }
// func deck(deck){
// 	arr:=[51]string{}
// 	a:=deck
// 	for i:=0;i<len(deck);i++{
// 		rand.Seed(time.Now().UnixNano())
// 		randomNum := random(0,51)
// 		temp:=a[randomNum]
// 		a[randomNum]=a[i]
// 		a[i]=temp
// 	}
// 	n:=0
// 	for i:=0;i<4;i++{
// 		for j:=0;j<9;j++{
// 			arr[i][j]=a[n]
// 			n++
// 		}
// 	}
// 	fmt.Println(arr)
// 	return arr
// }