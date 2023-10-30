package main
import "fmt"

func createMatrix(rows, cols int ) [][]int{
	matrix := [][]int{}
	for i := 0; i < rows; i++{
		row := []int{}
		for j := 0 ; j < cols ; j++{
			row = append(row,i*j)
		}
		matrix = append(matrix,row)
	}
	return matrix
}

func main(){
	m1 := createMatrix(4,4)
	for i := 0; i < len(m1);i++{
		fmt.Println(m1[i])
	}
}