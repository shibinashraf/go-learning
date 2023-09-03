package main
import "fmt"
func main(){
	num1 := 10
	num2 := 20
	if num1 > num2 {
		fmt.Println(num1, " is greater ", num2)
	} else {
		fmt.Println(num2, " is greater than ", num1)
	}
}