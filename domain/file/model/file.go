package model

import (
	"sync"

	"github.com/hanwen/go-fuse/fuse"
)

type openedFileServer struct {
	m              sync.Mutex
	openedFilesMap map[uint64]*Dentry
}

type Dentry struct {
	descriptor uint64
}

type DirEntry struct {
	Entries []*(fuse.DirEntry)
	Count   uint64
	m       sync.Mutex
}

func (d *DirEntry) RetrieveOneEntry() *fuse.DirEntry {
	d.m.Lock()
	defer d.m.Unlock()
	if len(d.Entries) == 0 {
		return nil
	}
	ret := d.Entries[len(d.Entries)-1]
	d.Entries = d.Entries[:len(d.Entries)-1]
	d.Count = d.Count - 1

	return ret
}

func (d *DirEntry) AddOneEntry(dentry *fuse.DirEntry) {
	d.m.Lock()
	defer d.m.Unlock()
	d.Entries = append(d.Entries, dentry)
	d.Count = d.Count + 1
}
