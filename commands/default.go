package commands

import (
	"fmt"
	"github.com/defeated/going/matcher"
	"github.com/defeated/going/storage"
)

func CmdDefault(input string, stor *storage.Storage) {
	keys := make([]string, 0, len(stor.Paths))
	for k := range stor.Paths {
		keys = append(keys, k)
	}
	find(input, keys)
}

func find(input string, items []string) {
	matches := matcher.Match(input, items)
	if len(matches) > 0 {
		directory := matches[0]
		fmt.Println(directory)
	}
}
