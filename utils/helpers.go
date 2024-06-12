package main

import "os"

// Helper function to open file or panic
func MustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
