package main

import (
	"fmt"
)

type greaterThanFifty struct {
	number int
}

func (gtf greaterThanFifty) Error() string {
	return fmt.Sprintf("%v Greater than 50!", gtf.number)
}

func checkGreater(num int) (int, error) {
	if num > 50 {
		return 0, greaterThanFifty{number: num}
	}
	return num, nil
}

func test(num int) {
	result, err := checkGreater(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v OK ",result)
	}
}

func main() {
	test(10)
	test(55)
}
