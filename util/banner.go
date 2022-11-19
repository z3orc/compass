package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner() {
	banner := figure.NewColorFigure("DynamicRPC", "smslant", "cyan", true)
	banner.Print()
	fmt.Println("----------------------------------------------------------------")
}