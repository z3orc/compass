package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner(text string) {
	banner := figure.NewColorFigure(text, "smslant", "cyan", true)
	banner.Print()
	fmt.Println("--------------------------------------------------v1.0.0-------")
}