package main
import (
	"fmt"
)


func findSum(nums ...int) int{
	sum := 0
	for i := 0; i < len(nums); i++{
		sum += nums[i]
	}
	return sum
}
func main(){
	total1 := findSum(1,2,3)
	fmt.Println(total1)
	total2 := findSum(1,2,3,4,5)
	fmt.Println(total2)

}