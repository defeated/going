package main

import (
	"fmt"
	"regexp"
	"github.com/defeated/going/reader"
)

func main() {
	lines := reader.Read("data.json")
	matcher := regexp.MustCompile("(?i).*w.*")
	matches := []string{}

	for _, dir := range d.Paths {
		match := matcher.MatchString(dir)
		if match {
			matches = append(matches, dir)
		}
	}

	fmt.Println(matches)
}
