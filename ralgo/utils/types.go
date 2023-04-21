package utils

type KeyCouple struct {
	A string
	B string
}

type Flags struct {
	Encode     bool
	Decode     bool
	InputText  string
	OutputFile string
	Couple     string
}

type EncoderInitializedData struct {
	Characters [][]int
	Depth      int
	Bits       int
	Key        string
}

type DecoderInitializedData struct {
	Chunks [][]string
	Depth  int
	Bits   int
}
