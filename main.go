package main

import (
	"fmt"

	"github.com/ZayaSansanch/Game/render"
)

func main() {
	fmt.Print("\033[H\033[2J")

	render.Render()
	// Render()
}
