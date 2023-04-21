package utils

import (
	"fmt"
	"os"
	"strings"
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

	if strings.Contains("01", keyCouple.A) || strings.Contains("01", keyCouple.B) {
		_, _ = fmt.Fprintln(os.Stderr, "Error: Key couple not contains 0 or 1")
		os.Exit(1)
	}
}

func CheckEncodedDataValidity(data []byte, couple KeyCouple) {
	for _, c := range data {
		if string(c) != couple.A && string(c) != couple.B {
			_, _ = fmt.Fprintln(os.Stderr, "Error: Encoded string must composed of only 2 characters of the couple")
			os.Exit(1)
		}
	}
}

func CheckError(err error, message ...string) {
	if err != nil {
		if message != nil {
			_, _ = fmt.Fprintln(os.Stderr, message)
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
