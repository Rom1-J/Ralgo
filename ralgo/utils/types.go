package utils

type KeyCouple struct {
	A string
	B string
}

type Flags struct {
	Encode     bool
	Decode     bool
	InputData  []byte
	OutputFile string
	Couple     string
}

type EncoderInitializedData struct {
	Characters [][]int64
	Depth      int
	Bits       int
	Key        string
}

type DecoderInitializedData struct {
	Chunks [][]byte
	Depth  int
	Bits   int
}
