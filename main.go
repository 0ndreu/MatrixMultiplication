package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 10
	matrix1 := make([][]int, n)
	matrix2 := make([][]int, n)
	matrix3 := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix1[i] = make([]int, n)
		matrix2[i] = make([]int, n)
		matrix3[i] = make([]int, n)
	}
	fillMatrix(matrix1)
	fillMatrix(matrix2)
	fmt.Println(matrix1)
	fmt.Println(matrix2)

}

func fillMatrix(matrix [][]int) {
	for i, innerArray := range matrix {
		for j := range innerArray {
			matrix[i][j] = rand.Intn(100)
			//looping over each element of array and assigning it a random variable
		}
	}
}
