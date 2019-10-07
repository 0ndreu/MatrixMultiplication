package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 2
	matrix1 := make([][]int, n)
	matrix2 := make([][]int, n)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix1[i] = make([]int, n)
		matrix2[i] = make([]int, n)
		sum[i] = make([]int, n)
	}
	fillMatrix(matrix1)
	fillMatrix(matrix2)
	for i := 0; i < n; i++ {
		go multiple(matrix1, matrix2, sum, n, i)
	}
	fmt.Println(matrix1)
	fmt.Println(matrix2)
	fmt.Println(sum)

}

func fillMatrix(matrix [][]int) {
	for i, innerArray := range matrix {
		for j := range innerArray {
			matrix[i][j] = rand.Intn(100)
			//looping over each element of array and assigning it a random variable
		}
	}
}

func multiple(matrix1, matrix2, sum [][]int, n int, i int) {
	for j := 0; j < n; j++ {
		sum[i][j] += matrix1[i][j] * matrix2[j][i]
	}
}
