package main

import (
	"fmt"
	"github.com/defeated/going/matcher"
	"github.com/defeated/going/reader"
	"os"
)

func main() {
	lines := reader.Read("data.json")
	input := os.Args[1]
	matches := matcher.Match(input, lines)
	directory := matches[0]

	fmt.Println(directory)
}
