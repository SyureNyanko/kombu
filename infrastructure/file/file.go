package file

import (
	_ "errors"
	"fmt"
	"sync"

	"github.com/hanwen/go-fuse/fuse"
)

type OpenedFileServerImpl struct {
	m              sync.Mutex
	openedFilesMap map[uint64]*(ImplDirEntry)
	count          uint64
}

/*
type OpenedFileServer interface {
	Register(f *model.Dentry) (uint64, error)
	Retrieve(d uint64)(*model.Dentry, error)
	Forget(d uint64) error
}

*/
/*
 type openedFileServer struct {
	m sync.Mutex
	openedFilesMap map[uint64]*Dentry
}
*/

func NewOpenedFileServer() *OpenedFileServerImpl {
	mp := map[uint64]*(ImplDirEntry){}
	return &OpenedFileServerImpl{count: 0, openedFilesMap: mp}
}

func (s *OpenedFileServerImpl) Register(f *ImplDirEntry) (uint64, error) {
	s.m.Lock()
	defer s.m.Unlock()

	var newcount uint64
	newcount = s.count + 1
	for {
		newcount = newcount + 1
		_, exist := s.openedFilesMap[newcount]
		if !exist {
			break
		}
	}
	s.openedFilesMap[newcount] = f
	return newcount, nil
}
func (s *OpenedFileServerImpl) Retrieve(d uint64) (*ImplDirEntry, error) {
	//s.m.lock()
	//defer s.m.unlock()
	v, _ := s.openedFilesMap[d]
	return v, nil
}

func (s *OpenedFileServerImpl) Forget(d uint64) error {
	s.m.Lock()
	defer s.m.Unlock()
	_, ok := s.openedFilesMap[d]
	if ok {
		delete(s.openedFilesMap, d)
		return nil
	}
	return nil
}

/*
type implDirEntry interface {
	RetrieveOneEntry() *fuse.DirEntryList
	AddOneEntry(dentry *fuse.DirEntryList)
}
*/

type ImplDirEntry struct {
	Entries []*(fuse.DirEntryList)
	Count   uint64
	m       sync.Mutex
}

func NewDirEntry() *ImplDirEntry {
	return &ImplDirEntry{
		Entries: make([]*(fuse.DirEntryList), 0),
		Count:   0,
	}
}


func (d *ImplDirEntry) RetrieveOneEntry() *fuse.DirEntryList{
	d.m.Lock()
	defer d.m.Unlock()
	fmt.Println(d.Count)
	if len(d.Entries) == 0 {
		return nil
	}
	ret :=  d.Entries[len(d.Entries)-1]
	d.Entries = d.Entries[:len(d.Entries)-1]
	d.Count = d.Count - 1;
	return ret
}

func (d *ImplDirEntry) AddOneEntry(dentry *fuse.DirEntryList) {
	d.m.Lock()
	defer d.m.Unlock()
	d.Entries = append(d.Entries, dentry)
	d.Count = d.Count + 1
}

/*
	Register(f *File) (uint64, error)
	Forget(d uint64) error
*/
