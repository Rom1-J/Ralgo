package utils

import "fmt"

func PrintMatrix(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func FlattenMatrix(matrix [][]int) []int {
	var flatMat []int

	for _, row := range matrix {
		flatMat = append(flatMat, row...)
	}

	return flatMat
}

func TransposeMatrix(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])

	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}
