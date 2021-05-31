package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Database struct {
	memory map[string][]byte
	file *File
}

type Entry struct {
	Id string `json:"id"'`
	Content []byte `json:"content"`
}

func NewDatabase() (*Database, error) {
	file := &File{path: "db"}
	memory, err := load(file)
	if err != nil {return nil, err}

	return &Database{
		memory: memory,
		file:   file,
	}, nil
}

func (d *Database) Get(id string) ([]byte, error) {
	result, ok := d.memory[id]
	if !ok {
		return nil, fmt.Errorf("item with id %s does not exist", id)
	}

	return result, nil
}

func (d *Database) Set(id string, content []byte) error {
	entry := Entry{
		Id: id,
		Content: content,
	}
	d.memory[id] = content

	b, err := json.Marshal(&entry)
	if err != nil {return err}

	s := string(b)
	line := s + "\n"

	return d.file.Write([]byte(line))
}

func load(file *File) (map[string][]byte, error) {
	db := make(map[string][]byte)

	b, err := file.Read()
	if err != nil {return nil, err}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		var entry Entry
		err := json.Unmarshal([]byte(line), &entry)
		if err != nil {continue}
		db[entry.Id] = entry.Content
	}

	return db, nil
}
