package model

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

type dataServer struct {
	id      string
	backend BackendType
}
