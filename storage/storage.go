package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type Path struct {
	Hits  uint16
	First time.Time
}

type Paths map[string]Path

type Storage struct {
	Filename string
	Paths    Paths
}

func NewStorage(filename string) *Storage {
	setup(filename)

	s := &Storage{filename, make(Paths)}
	s.Read()
	return s
}

func (s *Storage) Read() {
	f, _ := ioutil.ReadFile(s.Filename)
	json.Unmarshal(f, &s.Paths)
}

func (s *Storage) Write() {
	j, _ := json.Marshal(s.Paths)
	ioutil.WriteFile(s.Filename, j, 0666)
}

func setup(filename string) {
	dir := path.Dir(filename)
	os.MkdirAll(dir, 0700)
}
