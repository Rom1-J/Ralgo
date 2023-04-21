package ralgo

import (
	"fmt"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
	"math/rand"
	"strconv"
	"strings"
)

func EncoderInitializeData(message string, couple utils.KeyCouple) utils.EncoderInitializedData {
	var characters [][]int

	depth := 0

	highestFactor := 0
	bits := 0

	for _, char := range message {
		factors := utils.PrimeFactors(int(char))

		if len(factors) > depth {
			depth = len(factors)
		}

		for _, factor := range factors {
			if factor > highestFactor {
				highestFactor = factor
			}
		}

		characters = append(characters, factors)
	}

	bits = len(strconv.FormatInt(int64(highestFactor), 2))

	for i := range characters {
		for len(characters[i]) < depth {
			characters[i] = append(characters[i], 1)
		}
	}

	key := strings.Repeat(couple.A, depth) + strings.Repeat(couple.B, bits) + strings.Repeat(couple.A, bits) + couple.B

	return utils.EncoderInitializedData{
		Characters: characters,
		Depth:      depth,
		Bits:       bits,
		Key:        key,
	}
}

func Encode(message string, couple utils.KeyCouple) {
	initializedData := EncoderInitializeData(message, couple)

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
