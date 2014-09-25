package main

import (
	"fmt"

	"github.com/monochromegane/cgotest"
)

func main() {
	fmt.Printf("HOME: %s\n", cgotest.GetHome())
}
