package model

import "os"

type BackendType int



const (
	LocalFileSystem BackendType = iota
)

func (c BackendType) String() string {
	switch c {
	case LocalFileSystem:
		return "LocalFileSystem"
	default:
		return "Unknown"
	}
}

type DataServer struct {
	Id            string
	Backend       BackendType
}

type DataFile struct {
	fd *os.File /* need when read data from file */
}

type Writer struct{}

type Reader struct{}