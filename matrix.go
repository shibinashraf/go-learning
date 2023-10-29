package main
import "fmt"

func createMatrix(rows, cols int) [][]int{
	matrix := [][]int{}
	for i := 0 ; i < rows ; i++{
		row := []int{}
		for j :=0; j < cols ; j++{
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}


func main(){
	matrix := createMatrix(4,4)
	for i := 0; i< len(matrix); i++{
		fmt.Println(matrix[i])
	}
}