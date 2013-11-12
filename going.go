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

	matcher := regexp.MustCompile("(?i).*b.*")

	for _, dir := range dirs {
		match := matcher.MatchString(dir)

		if match {
			fmt.Println(dir)
		}
	}
}
