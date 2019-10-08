package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var n int
	fmt.Print("Введите размер квадратной матрицы ")
	fmt.Scan(&n)
	matrix1 := make([][]int, n)
	matrix2 := make([][]int, n)
	sum := make([][]int, n)
	sumWithoutRoutine := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix1[i] = make([]int, n)
		matrix2[i] = make([]int, n)
		sum[i] = make([]int, n)
		sumWithoutRoutine[i] = make([]int, n)
	}
	fillMatrix(matrix1)
	fillMatrix(matrix2)

	t1 := time.Now()
	for i := 0; i < n; i++ {
		go multiple(matrix1, matrix2, sum, n, i)
	}
	t2 := time.Now()
	fmt.Println(int(t2.Sub(t1)/time.Microsecond), " microseconds with goroutine")

	t3 := time.Now()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				sum[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	t4 := time.Now()
	fmt.Println(int(t4.Sub(t3)/time.Microsecond), " microseconds without goroutine")

	// fmt.Println(matrix1)
	// fmt.Println(matrix2)
	// fmt.Println(sum)
	if n < 4 {
		outputMatrix(matrix1, n)
		fmt.Print("\n\n")
		outputMatrix(matrix2, n)
		fmt.Print("\n\n")
		outputMatrix(sum, n)
	}
	fmt.Println("Correct")

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
		for k := 0; k < n; k++ {
			sum[i][j] += matrix1[i][k] * matrix2[k][j]
		}
	}
}

func outputMatrix(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
