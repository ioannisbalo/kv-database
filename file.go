package main

import "os"

type File struct {
	path string
}

func (f *File) Write(line []byte) error {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {return err}
	defer file.Close()

	_, err = file.Write(line)
	return err
}

func (f *File) Read() ([]byte, error) {
	file, err := os.OpenFile(f.path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {return nil, err}
	defer file.Close()

	b := make([]byte, 1024)
	_, err = file.Read(b)

	return b, nil
}
