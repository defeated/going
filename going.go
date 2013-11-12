package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

type Data struct {
	Paths []string
}

func main() {
	matcher := regexp.MustCompile("(?i).*w.*")
	matches := []string{}

	var d Data
	f, _ := ioutil.ReadFile("data.json")
	json.Unmarshal(f, &d)

	for _, dir := range d.Paths {
		match := matcher.MatchString(dir)
		if match {
			matches = append(matches, dir)
		}
	}

	fmt.Println(matches)
}
