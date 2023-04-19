package utils

type KeyCouple struct {
	A string
	B string
}

type Flags struct {
	InputText  string
	OutputFile string
}

type InitializedData struct {
	Characters [][]int
	Depth      int
	Bits       int
}
