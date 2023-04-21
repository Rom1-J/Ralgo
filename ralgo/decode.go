package ralgo

import (
	"fmt"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
)

func DecoderInitializeData(message string, couple utils.KeyCouple) utils.DecoderInitializedData {
	utils.CheckEncodedMessageValidity(message, couple)

	var chunks [][]string

	depth := 0
	bits := 0

	return utils.DecoderInitializedData{
		Chunks: chunks,
		Depth:  depth,
		Bits:   bits,
	}
}

func Decode(message string, couple utils.KeyCouple) {
	initializedData := DecoderInitializeData(message, couple)

	fmt.Printf("bits: %d\n", initializedData.Bits)
	fmt.Printf("depth: %d\n", initializedData.Depth)

	fmt.Println()
}
