package main 
import "fmt"


func printPrimes(max int ) {
	for i := 2 ; i < max+1 ; i++{
		if i==2{
			fmt.Println(i)
			continue
		}
		if i%2 ==0{
			continue
		}
		for j :=3; j*j < i ; j++ {
			if j%i == 0{
				continue
			}
		}
		fmt.Println(i)
	}
}
func main(){

	printPrimes(100)
}