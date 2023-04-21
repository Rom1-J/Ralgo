package utils

import (
	"flag"
)

func GetFlags() Flags {
	var encode bool
	var decode bool

	flag.BoolVar(&encode, "e", false, "Encode text")
	flag.BoolVar(&decode, "d", false, "Decode text")

	inputText := flag.String("input", "", "Input text")
	outputFile := flag.String("output", "", "Output file (default \"stdout\")")

	couple := flag.String("couple", ".*", "Two characters to use has couple")

	flag.Parse()

	return Flags{
		Encode:     encode,
		Decode:     decode,
		InputText:  *inputText,
		OutputFile: *outputFile,
		Couple:     *couple,
	}
}
