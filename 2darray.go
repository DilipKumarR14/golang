package main
import "fmt"
func main(){
	var sizes int 
	var ranges int
	var low=0
	var high=ranges
	fmt.Println("Enter the range of numbers : ")
	fmt.Scanf("%d",&sizes)
	fmt.Println("enter the no of element to be stored in each row : ")
	fmt.Scanf("%d",&ranges)	
	var result=sizes/ranges
	// arr:=[][] int{}
	arr := make([][]int,sizes) // declare the 2d array
	for i := range arr {
		arr[i] = make([]int, sizes)
	}
	var k=0
	for i:=0;i<result;i++{
		if i==0 {
			low=low
			high=high
		}else{
			low=high+1
			high=high+ranges
		}
		for j := low; j <= high; j++ {
			if prime(j){
				arr[i][k]=j
				k++
			}
		}
	}
	fmt.Println(arr)
}
func prime(size int) bool {
	var flag =true

	for i:=1;i<size;i++{
		for j:=2;j<i;j++{
			if i%j==0{
				flag=false
				break
			}
		}
		if flag==true{
			return true
		}
	}
	return flag
}