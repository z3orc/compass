package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner(text string) {
	banner := figure.NewColorFigure(text, "smslant", "purple", true)
	banner.Print()
	fmt.Println("--------------------------------------------------v2.0.0-------")
}
