package main

import (
	"./commands"
	"os"
)

func main() {
	cmd := os.Args[1]
	if cmd == "--add" {
		commands.CmdAdd()
	} else {
		commands.CmdDefault(cmd, "data.json")
	}
}
