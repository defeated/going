package commands

import (
	"../matcher"
	"../reader"
	"fmt"
)

func CmdDefault(input string, filename string) {
	lines := reader.Read(filename)
	matches := matcher.Match(input, lines)

	if len(matches) > 0 {
		directory := matches[0]
		fmt.Println(directory)
	}
}
