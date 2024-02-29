package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner(text string, version string, build string) {
	banner := figure.NewColorFigure(text, "smslant", "cyan", true)
	banner.Print()
	fmt.Printf("-----------------------v%s--------------\n", version)
}
