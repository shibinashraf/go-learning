package main

import ( 
	"fmt"
)

func main(){

	items := []string{"apple","car", "mobile"}
	items = append(items, "laptop")
	fmt.Printf("Item 1 : %v \n", items[0])
	fmt.Printf("Item 2 : %v \n", items[1])
	fmt.Printf("Item 3 : %v \n", items[2])
	fmt.Printf("Item 4 : %v \n", items[3])
	fmt.Printf("some of the items are : %v ", items[0:2])

}