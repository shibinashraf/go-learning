package main
import "fmt"


func main(){

	type car struct{
		model string
		color string
		topSpeed int
	}
	
	carOne := car{
		model : "Tesla",
		color: "red",
		topSpeed : 120,
	}

	fmt.Printf("The car is %s in %s color and has a top speed of %d",carOne.model, carOne.color, carOne.topSpeed)
}