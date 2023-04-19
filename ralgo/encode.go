package ralgo

import (
	"fmt"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
	"math/rand"
	"strconv"
	"strings"
)

func Encode(message string, couple utils.KeyCouple) {
	initializedData := utils.EncoderInitializeData(message, couple)

	matrix := initializedData.Characters

	fmt.Printf("bits: %d\n", initializedData.Bits)
	fmt.Printf("depth: %d\n", initializedData.Depth)
	fmt.Printf("key: %s\n", initializedData.Key)

	fmt.Println()

	rows, _ := initializedData.Depth, len(matrix)
	for _, col := range matrix {
		rand.Shuffle(rows, func(i, j int) {
			col[i], col[j] = col[j], col[i]
		})
	}

	matrix = utils.TransposeMatrix(matrix)

	output := initializedData.Key

	for _, el := range utils.FlattenMatrix(matrix) {
		binaryEl := strconv.FormatInt(int64(el), 2)

		for _, char := range strings.Repeat("0", initializedData.Bits-len(binaryEl)) + binaryEl {
			if char == 48 {
				output += couple.A
			} else {
				output += couple.B
			}
		}
	}

	fmt.Println(output)
}
