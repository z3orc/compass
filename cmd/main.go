package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/z3orc/compass/internal/data"
	"github.com/z3orc/compass/internal/repo"
)

func main() {
	log.SetLevel(log.DebugLevel)

	fmt.Println("Hello World!")
	src := data.NewPistonDataSource()
	src.GetVersion("1.21")

	repo.NewVersionRepository(src)
}
