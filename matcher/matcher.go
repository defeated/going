package matcher

import (
	"regexp"
	"strings"
)

func fuzzify(input string) string {
	chars := strings.Split(input, "")
	fuzzy := ".*" + strings.Join(chars, ".*")
	return fuzzy
}

func Match(input string, list []string) []string {
	pattern := regexp.MustCompile("(?i)" + fuzzify(input))
	matches := []string{}

	for _, dir := range list {
		match := pattern.MatchString(dir)
		if match {
			matches = append(matches, dir)
		}
	}

	return matches
}
