package storage

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Path struct {
	Hits  uint16
	First time.Time
}

type Storage struct {
	Filename string
	Paths    map[string]Path
}

func NewStorage(filename string) *Storage {
	s := &Storage{Filename: filename}
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
