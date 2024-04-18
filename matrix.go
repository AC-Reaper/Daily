package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

type ElementResult struct {
	row   int
	col   int
	value int
}

func multiplyRowByColumn(a [][]int, b [][]int, row, col int, resultChan chan<- ElementResult) {
	// add your code
	value := 0
	for k := 0; k < len(a); k++ {
		value += a[row][k] * b[k][col]
	}
	resultChan <- ElementResult{
								row : row,
								col : col,
								value : value}
}

func concurrentMatrixMultiply(a, b [][]int) [][]int {
	numRows, numCols := len(a), len(b[0])
	result := make([][]int, numRows)
	for i := range result {
		result[i] = make([]int, numCols)
	}

	resultChan := make(chan ElementResult, numRows*numCols)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			go multiplyRowByColumn(a, b, i, j, resultChan)
		}
	}

	for i := 0; i < numRows*numCols; i++ {
		res := <-resultChan
		result[res.row][res.col] = res.value
	}
	close(resultChan)

	return result
}

func generateMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = 1
		}
	}
	return matrix
}

func main() {

	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	var size = flag.Int("s", 1000, "size of the matrix")

	flag.Parse()
	runtime.GOMAXPROCS(*numCores)

	a := generateMatrix(*size, *size)
	b := generateMatrix(*size, *size)

	startTime := time.Now()

	result := concurrentMatrixMultiply(a, b)

	duration := time.Since(startTime)

	fmt.Println("result[0][0] = ", result[0][0])

	fmt.Printf("Matrix multiplication completed. Size: %dx%d, Time taken: %v\n", *size, *size, duration)

}
