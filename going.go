package main

import (
	"github.com/defeated/going/commands"
	"github.com/defeated/going/storage"
	"os"
)

func main() {
	stor := storage.NewStorage("data.json")
	cmd := os.Args[1]
	if cmd == "--add" {
		path := os.Args[2]
		commands.CmdAdd(stor, path)
	} else {
		commands.CmdDefault(stor, cmd)
	}
}
