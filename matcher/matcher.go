package matcher

import "regexp"

func Match(input string, list []string) []string {
	pattern := regexp.MustCompile("(?i).*" + input + ".*")
	matches := []string{}

	for _, dir := range list {
		match := pattern.MatchString(dir)
		if match {
			matches = append(matches, dir)
		}
	}

	return matches
}
