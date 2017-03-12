package main

import (
	"fmt"

	"github.com/raeoks/malinche"
)

func main() {
	version := fmt.Sprintf(
		"Malinche Version %s: %s",
		malinche.Version,
		malinche.Build)
	fmt.Println(version)
}
