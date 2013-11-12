package reader

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	Paths []string
}

func Read(filename string) []string {
	var d Data
	f, _ := ioutil.ReadFile(filename)
	json.Unmarshal(f, &d)
	return d.Paths
}
