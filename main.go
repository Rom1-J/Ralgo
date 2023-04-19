package main

import (
	"github.com/Rom1-J/Ralgo/ralgo"
	"github.com/Rom1-J/Ralgo/ralgo/utils"
)

func main() {
	flags := utils.GetFlags()
	utils.CheckFlags(flags)

	ralgo.Encode(flags.InputText)
}
