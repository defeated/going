package commands

import (
	"github.com/defeated/going/storage"
	"time"
)

func CmdAdd(stor *storage.Storage, key string) {
	path, ok := stor.Paths[key]
	if ok {
		path.Hits++
	} else {
		path = storage.Path{1, time.Now().UTC()}
	}

	stor.Paths[key] = path
	stor.Write()
}
