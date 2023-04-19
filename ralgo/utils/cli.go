package utils

import (
	"flag"
)

func GetFlags() Flags {
	inputText := flag.String("input", "", "Input text")
	outputFile := flag.String("output", "", "Output file (default \"stdout\")")

	flag.Parse()

	return Flags{
		InputText:  *inputText,
		OutputFile: *outputFile,
	}
}
