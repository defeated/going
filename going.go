package main

import (
	"fmt"
	"regexp"
)

func main() {
	dirs := []string{
		"aaa",
		"bbb",
	}

	for _, dir := range dirs {
		match, _ := regexp.MatchString("(?i).*b.*", dir)

		if match {
			fmt.Println(dir)
		}
	}
}
