package commands

import (
	"fmt"
	"github.com/defeated/going/matcher"
	"github.com/defeated/going/storage"
)

func CmdDefault(stor *storage.Storage, input string) {
	findMatch(input, getKeys(stor.Paths))
}

func getKeys(hash storage.Paths) []string {
	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	return keys
}

func findMatch(input string, items []string) {
	matches := matcher.Match(input, items)
	if len(matches) > 0 {
		directory := matches[0]
		fmt.Println(directory)
	}
}
