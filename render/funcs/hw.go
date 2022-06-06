package funcs

import (
	"golang.org/x/term"
)

func Hw() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	return height, width
}
