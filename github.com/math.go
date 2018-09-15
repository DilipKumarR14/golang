package main
import "fmt"
// import "github.com/utility"
func main(){
	// utility.Add()
	var size int
	fmt.Println("enter the size : ")
	fmt.Scanf("%d",&size)
 array:=make([] int,size)
 for i:=0;i<len(array);i++{
	 fmt.Println("enter the elements")
	 fmt.Scanf("%d",&array[i])
 }
 fmt.Println(array)
 for i:=0;i<len(array);i++{
	 for j:=i;j<len(array);j++{
		 if array[i]>array[j]{
			 array[i],array[j]=array[j],array[i]
		 }
	 }
 }
 
 fmt.Println(array)
}  