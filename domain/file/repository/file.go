package repository

import (
	"github.com/hanwen/go-fuse/fuse"
)

type OpenedFileServer interface {
	Register(DirEntry) (uint64, error)
	Retrieve(uint64)(DirEntry, error)
	Forget(uint64) error
}


type DirEntry interface {
	RetrieveOneEntry() *fuse.DirEntryList
	AddOneEntry(*fuse.DirEntryList)
}