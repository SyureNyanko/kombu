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

type DataServer struct {
	id      string
	backend BackendType
}

/*
 boundaru
*/
type data struct {
	uri string
}
