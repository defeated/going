package main

import (
	"github.com/defeated/going/commands"
	"github.com/defeated/going/storage"
	"os"
)

func main() {
	stor := storage.NewStorage(os.Getenv("HOME") + "/.going/data.json")
	cmd := os.Args[1]
	switch cmd {
	default:
		commands.CmdDefault(stor, cmd)
	case "--add":
		path := os.Args[2]
		commands.CmdAdd(stor, path)
	case "--prune":
		commands.CmdPrune(stor)
	}
}
