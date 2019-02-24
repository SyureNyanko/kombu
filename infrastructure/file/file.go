package file

import (
	_ "errors"
	_ "fmt"
	"sync"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/domain/file/model"
)

type OpenedFileServerImpl struct {
	m              sync.Mutex
	openedFilesMap map[uint64]*(model.DirEntry)
	count          uint64
}

var sharedInstance *OpenedFileServerImpl = newOpenedFileServer()

func NewOpenedFileServer() *OpenedFileServerImpl {
	return sharedInstance
}

func newOpenedFileServer() *OpenedFileServerImpl {
	mp := map[uint64]*(model.DirEntry){}
	return &OpenedFileServerImpl{count: 0, openedFilesMap: mp}
}

func (s *OpenedFileServerImpl) Register(f *model.DirEntry) (uint64, error) {
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
func (s *OpenedFileServerImpl) Retrieve(d uint64) (*model.DirEntry, error) {
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

func (s *OpenedFileServerImpl) NewDirEntry() *model.DirEntry {
	return &model.DirEntry{
		Entries: make([]*(fuse.DirEntry), 0),
		Count:   0,
	}
}
