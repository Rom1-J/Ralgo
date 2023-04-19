package utils

import (
	"fmt"
	"os"
)

func CheckFlags(flags Flags) {}

func CheckKeyCouple(couple KeyCouple) {
	if couple.A == couple.B {
		fmt.Println("Key couple must be 2 distincts characters")
		os.Exit(1)
	}
}
