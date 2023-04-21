package ralgo

import (
	"fmt"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
	"math/rand"
	"strconv"
	"strings"
)

func EncoderInitializeData(data []byte, couple utils.KeyCouple) utils.EncoderInitializedData {
	var bytes [][]int64

	factorMap := map[byte][]int64{
		0: {0},
		1: {1},
	}

	depth := 0

	highestFactor := int64(0)
	bits := 0

	for _, byteValue := range data {
		factors := []int64{0}

		if f, ok := factorMap[byteValue]; ok {
			factors = f
		} else {
			f = utils.PrimeFactors(int(byteValue))
			factorMap[byteValue] = f
			factors = f
		}

		if len(factors) > depth {
			depth = len(factors)
		}

		lastFactor := factors[len(factors)-1]
		if lastFactor > highestFactor {
			highestFactor = lastFactor
		}

		bytes = append(bytes, factors)
	}

	bits = len(strconv.FormatInt(highestFactor, 2))

	for i := range bytes {
		for len(bytes[i]) < depth {
			bytes[i] = append(bytes[i], 1)
		}
	}

	key := strings.Repeat(couple.A, depth) + strings.Repeat(couple.B, bits) + couple.A

	return utils.EncoderInitializedData{
		Characters: bytes,
		Depth:      depth,
		Bits:       bits,
		Key:        key,
	}
}

func Encode(data []byte, couple utils.KeyCouple) {
	initializedData := EncoderInitializeData(data, couple)

	matrix := initializedData.Characters

	cols := initializedData.Depth

	for _, row := range matrix {
		rand.Shuffle(cols, func(i, j int) {
			row[i], row[j] = row[j], row[i]
		})
	}

	matrix = utils.TransposeMatrix(matrix)

	fmt.Print(initializedData.Key)

	for _, el := range utils.FlattenMatrix(matrix) {
		binaryEl := strconv.FormatInt(el, 2)

		fmt.Print(
			strings.Replace(
				strings.Replace(
					strings.Repeat("0", initializedData.Bits-len(binaryEl))+binaryEl,
					"0",
					couple.A,
					-1,
				),
				"1",
				couple.B,
				-1,
			),
		)
	}
}
