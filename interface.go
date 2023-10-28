package main
import "fmt"

type Printer interface{
	printInfo() 
}

type drink struct{
	name  string
	price  float64
	dType  string
}

type food struct{
	name  string
	price  float64
}

func typeChecker( i interface{}){
	switch i.(type){
case string : fmt.Println("I am a string \n")
case int: fmt.Println("I am an int : %v \n", i)
case drink : fmt.Printf("I am a drink : %v \n ",i.(drink).name)
	}
}

func (d drink) printInfo() {
	fmt.Printf(
		"Drink Name : %s , Price : %.2f , Drink type : %s \n", d.name, d.price, d.dType)
}
func (f food) printInfo() {
	fmt.Printf(
		"Food Name : %s , Price : %.2f \n", f.name, f.price)
}

func main(){
	pizza := food{
		name : "Pizza",
		price : 12.5,
	}

	coffee := drink{
		name : "Coffee",
		price : 2.5,
		dType : "Hot",
	}

	coffee.printInfo()
	pizza.printInfo()
	

	typeChecker(coffee)
}