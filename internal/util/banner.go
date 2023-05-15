package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner(text string) {
	banner := figure.NewColorFigure(text, "smslant", "purple", true)
	version := "2.0.0"
	banner.Print()
	fmt.Printf("-----------------------------v%s--------\n", version)
}
