package main
import (
	"fmt"
	"errors"
)
func checkMessage(message , badwords []string) (bool, error){
	for _, word := range message{
		for _, badword := range(badwords){
			if word == badword{
				return false, errors.New("Not proper msg!")
			}
		}
	}
	return true,nil
}
func main(){

	message := []string{"hello","flick","go","kill"}
	badwords := []string{"flick","kill","die"}
	status, err := checkMessage(message, badwords)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(status)
	}
}