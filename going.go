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
		commands.CmdAdd()
	} else {
		commands.CmdDefault(cmd, stor)
	}
}
