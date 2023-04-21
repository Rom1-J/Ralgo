package utils

import (
	"fmt"
	"os"
)

func CheckFlags(flags Flags) {
	if flags.Encode == flags.Decode {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Either -e or -d must be specified, but not both")
		os.Exit(1)
	}

	if flags.Couple != ".*" {
		CheckKeyCouple(flags.Couple)
	}
}

func CheckKeyCouple(couple string) {
	if len(couple) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Please provide only two characters for the couple")
		os.Exit(1)
	}

	keyCouple := KeyCouple{
		A: string(couple[0]),
		B: string(couple[1]),
	}

	if keyCouple.A == keyCouple.B {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Key couple must be 2 distincts characters")
		os.Exit(1)
	}
}

func CheckEncodedMessageValidity(message string, couple KeyCouple) {
	for _, c := range message {
		if string(c) != couple.A && string(c) != couple.B {
			_, _ = fmt.Fprintln(os.Stderr, "Error: Encoded string must composed of only 2 characters of the couple")
			os.Exit(1)
		}
	}
}
