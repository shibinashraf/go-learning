package main
import "fmt"

func main(){
	fruits := []string{"apple","orange","watermelon","pineapple"}

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}