package commands

import (
	"github.com/defeated/going/storage"
)

func CmdPrune(stor *storage.Storage) {
	stor.Delete()
}
