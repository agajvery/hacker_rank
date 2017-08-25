package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var i, j int64
	var mainDiagonal, secondDiagonal int64
	var result int64

	matrix, matrixN := getInputParams()

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
	result = mainDiagonal - secondDiagonal
	fmt.Printf("%v", math.Abs(float64(result)))
}

func getInputParams() (matrix [][]string, matrixN int64) {
	reader := bufio.NewReader(os.Stdin)
	strInputMatrixN, _ := reader.ReadString('\n')
	matrixN, error := strconv.ParseInt(strings.TrimRight(strInputMatrixN, "\n"), 10, 32)

	if error == nil {
		for idx := 0; int64(idx) < matrixN; idx++ {
			strInputItems, _ := reader.ReadString('\n')
			rowMatrix := strings.Split(strings.TrimRight(strInputItems, "\n"), " ")
			matrix = append(matrix, rowMatrix)
		}
	}
	return
}

func getTestInputParams() (matrix [][]string, matrixN int64) {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	inputMatrixN := scanner.Text()

	matrixN, _ = strconv.ParseInt(inputMatrixN, 10, 64)

	for i := 0; int64(i) < matrixN; i++ {
		scanner.Scan()
		row := scanner.Text()
		rowMatrix := strings.Split(strings.TrimRight(row, "\n"), " ")
		matrix = append(matrix, rowMatrix)
	}
	return
}
