package util

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func Banner(text string, version string, build string) {
	banner := figure.NewColorFigure(text, "smslant", "purple", true)
	banner.Print()
	fmt.Printf("--------------------------v%s %s----\n", version, build)
}
