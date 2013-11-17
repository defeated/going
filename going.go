package main

import (
	"github.com/defeated/going/commands"
	"github.com/defeated/going/storage"
	"os"
)

func main() {
	stor := storage.NewStorage(os.Getenv("HOME") + "/.going/data.json")
	cmd := os.Args[1]
	if cmd == "--add" {
		path := os.Args[2]
		commands.CmdAdd(stor, path)
	} else if cmd == "--prune" {
		commands.CmdPrune(stor)
	} else {
		commands.CmdDefault(stor, cmd)
	}
}
