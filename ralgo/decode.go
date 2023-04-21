package ralgo

import (
	"github.com/Rom1-J/Ralgo/ralgo/utils"
	"os"
)

func DecoderInitializeData(data []byte, couple utils.KeyCouple) utils.DecoderInitializedData {
	utils.CheckEncodedDataValidity(data, couple)

	var chunks [][]byte

	depth := 0
	bits := 0

	step := 0
	lastChar := byte(0)
	var tmpChunk []byte

	for _, char := range data {
		if step == 2 {
			if len(tmpChunk) < bits {
				tmpChunk = append(tmpChunk, char)
			} else {
				chunks = append(chunks, tmpChunk)
				tmpChunk = append([]byte{}, char)
			}
		}
		if step == 1 {
			if lastChar == char {
				bits += 1
			} else {
				step += 1
			}
		}
		if step == 0 {
			if lastChar == 0 || lastChar == char {
				depth += 1
			} else {
				step += 1
				bits += 1
			}
		}
		lastChar = char
	}
	chunks = append(chunks, tmpChunk)

	return utils.DecoderInitializedData{
		Chunks: chunks,
		Depth:  depth,
		Bits:   bits,
	}
}

func Decode(data []byte, couple utils.KeyCouple) {
	initializedData := DecoderInitializeData(data, couple)

	rows, cols := initializedData.Depth, len(initializedData.Chunks)/initializedData.Depth

	matrix := make([][]int64, rows)
	for i := range matrix {
		matrix[i] = make([]int64, cols)
	}

	for i, chunk := range initializedData.Chunks {
		q, r := utils.EuclideanDivision(i, cols)
		matrix[q][r] = utils.CoupleToInt64(chunk, couple)
	}

	matrix = utils.TransposeMatrix(matrix)

	for _, chunk := range matrix {
		byteValue := int64(1)
		for _, factor := range chunk {
			byteValue *= factor
		}

		_, _ = os.Stdout.Write([]byte{byte(byteValue)})
	}
}
