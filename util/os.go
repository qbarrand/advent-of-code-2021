package util

import "os"

func MustOpen(name string) *os.File {
	fd, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return fd
}
