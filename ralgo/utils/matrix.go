package utils

import "fmt"

func PrintMatrix(matrix [][]int64) {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func FlattenMatrix(matrix [][]int64) []int64 {
	var flatMat []int64

	for _, row := range matrix {
		flatMat = append(flatMat, row...)
	}

	return flatMat
}

func TransposeMatrix(matrix [][]int64) [][]int64 {
	rows, cols := len(matrix), len(matrix[0])

	transposed := make([][]int64, cols)
	for i := range transposed {
		transposed[i] = make([]int64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
