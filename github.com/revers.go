package main
import "fmt"
import "sort"

func main(){
	num1 := []int {1,2,4,6,8,91}
	sort.Sort((sort.IntSlice(num1)))
	fmt.Println(num1)

	fmt.Println("Interger Reverse Sort")
    num := []int{50,90, 30, 10, 50}
    sort.Sort(sort.Reverse(sort.IntSlice(num)))
    fmt.Println(num)
}