package main

import (
	"fmt"

	"github.com/z3orc/compass/internal/data"
)

func main() {
	fmt.Println("Hello World!")
	src := data.NewPistonDataSource()
	src.GetVersion("1.21")
}
