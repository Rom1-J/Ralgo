package main

import (
	"github.com/Rom1-J/Ralgo/ralgo"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
)

func main() {
	flags := utils.GetFlags()
	utils.CheckFlags(flags)

	if flags.Encode {
		ralgo.Encode(flags.InputText, utils.KeyCouple{
			A: string(flags.Couple[0]),
			B: string(flags.Couple[1]),
		})
	}

	if flags.Decode {
		ralgo.Decode(flags.InputText, utils.KeyCouple{
			A: string(flags.Couple[0]),
			B: string(flags.Couple[1]),
		})
	}
}
