package model

import (
	"sync"
)

type openedFileServer struct {
	m              sync.Mutex
	openedFilesMap map[uint64]*Dentry
}

type Dentry struct {
	descriptor uint64
}

type DirEntry struct {
}
