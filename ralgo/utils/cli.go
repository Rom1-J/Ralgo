package utils

import (
	"flag"
	"os"
)

func getData(data string) []byte {
	content, err := os.ReadFile(data)
	if err != nil {
		return []byte(data)
	}

	return content
}

func GetFlags() Flags {
	var encode bool
	var decode bool

	flag.BoolVar(&encode, "e", false, "Encode text")
	flag.BoolVar(&decode, "d", false, "Decode text")

	inputData := flag.String("input", "", "Input text or file")
	outputFile := flag.String("output", "", "Output file (default \"stdout\")")

	couple := flag.String("couple", ".*", "Two characters to use has couple")

	flag.Parse()

	return Flags{
		Encode:     encode,
		Decode:     decode,
		InputData:  getData(*inputData),
		OutputFile: *outputFile,
		Couple:     *couple,
	}
}
