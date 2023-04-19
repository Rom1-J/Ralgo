package ralgo

import (
	"fmt"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
	"math/rand"
)

func Encode(message string) {
	initializedData := utils.InitializeData(message)

	matrix := initializedData.Characters

	rows, cols := initializedData.Depth, len(matrix)
	for _, col := range matrix {
		rand.Shuffle(rows, func(i, j int) {
			col[i], col[j] = col[j], col[i]
		})
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d ", matrix[j][i])
		}
		fmt.Println()
	}
}
