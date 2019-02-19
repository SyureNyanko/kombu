package repository

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/domain/file/model"
)

type OpenedFileServer interface {
	Register(*model.DirEntry) (uint64, error)
	Retrieve(uint64)(*model.DirEntry, error)
	Forget(uint64) error
	NewDirEntry() *model.DirEntry
}


type DirEntry interface {
	RetrieveOneEntry() *fuse.DirEntryList
	AddOneEntry(*fuse.DirEntryList)
}