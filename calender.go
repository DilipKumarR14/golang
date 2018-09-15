package main
import (
"fmt"
"math"
)

func GetCal() {
var month, year int32
fmt.Println("enter the month ")
fmt.Scanf("%d", &month)
fmt.Println("enter the year ")
fmt.Scanf("%d", &year)
months := []string{"", "January", "February", "March",
"April", "May", "June", "July", "August", "September",
"October", "November", "December"}
days := []int32{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

if month == 2 && LeapYear(year) {
days[month] = 29
}
fmt.Println(months[month], year)
fmt.Println("Su Mo Tu We Th Fr Sa")
d := Day(month, 1, year)
var i int32
for i = 0; i < d; i++ {
fmt.Print("   ")
}
for i = 1; i <= days[month]; i++ {
fmt.Printf("%2d ", i)
if ((i+d)%7 == 0) || (i == days[month]) {
fmt.Println()
}
}

}
func Day(month int32, day int32, year int32) int32 {

mon := float64(month)
yea := float64(year)
da := float64(day)
y := yea - math.Floor((14-mon)/12)
x := y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400)
m := mon + 12*math.Floor((14-mon)/12) - 2
d := (int32)(da+x+math.Floor((31*m)/12)) % 7
return d
}
func LeapYear(year int32) bool {
if (year%4 == 0) && (year%100 != 0) {
return true
}
if year%400 == 0 {
return true
}
return false
}

func main(){
	GetCal()
}
