package utils

import (
	"strconv"
)

func InitializeData(message string) InitializedData {
	var characters [][]int

	depth := 0

	highestFactor := 0
	bits := 0

	for _, char := range message {
		factors := PrimeFactors(int(char))

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

	return InitializedData{
		Characters: characters,
		Depth:      depth,
		Bits:       bits,
	}
}
