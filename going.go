package main

import (
	"fmt"
	"github.com/defeated/going/matcher"
	"github.com/defeated/going/reader"
)

func main() {
	lines := reader.Read("data.json")
	matches := matcher.Match("w", lines)
	fmt.Println(matches)
}
