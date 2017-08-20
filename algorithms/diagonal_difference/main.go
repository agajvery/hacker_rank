package main

import (
	"fmt"
	"strconv"
)

func main() {
	var matrix [][]string
	var i, j int64
	var mainDiagonal, secondDiagonal int64
	var result int64

	// reader := bufio.NewReader(os.Stdin)
	// strInputMatrixN, _ := reader.ReadString('\n')
	// matrixN, error := strconv.ParseInt(strings.TrimRight(strInputMatrixN, "\n"), 10, 32)

	// if error == nil {
	// 	for idx := 0; int64(idx) < matrixN; idx++ {
	// 		strInputItems, _ := reader.ReadString('\n')
	// 		rowMatrix := strings.Split(strings.TrimRight(strInputItems, "\n"), " ")
	// 		matrix = append(matrix, rowMatrix)
	// 	}
	// }
	matrixN := 5
	matrix = append(matrix, {"-10", "3", "0", "5", "4"})
	for ; i < matrixN; i++ {
		for j = 0; j < matrixN; j++ {
			if i == j {
				value, _ := strconv.ParseInt(matrix[i][j], 10, 64)
				mainDiagonal += value
			}
			if (i + j) == matrixN-1 {
				value, _ := strconv.ParseInt(matrix[i][j], 10, 64)
				secondDiagonal += value
			}
		}
	}
	result = (mainDiagonal - secondDiagonal) * -1
	fmt.Printf("%v", result)
}
