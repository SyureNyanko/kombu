package inode

import (
	"io/ioutil"
	"strconv"
	"sync"

	"github.com/kombu/domain/repository"
)

type InodeServerImpl struct {
	filepath string
	mu       sync.Mutex
}

func NewInodeServerImpl(inodefilepath string) repository.InodeServer {
	return &InodeServerImpl{filepath: "maxinode"}
}

func (s *InodeServerImpl) IssueId() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := ioutil.ReadFile(s.filepath)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	ioutil.WriteFile(s.filepath, []byte(strconv.Itoa(i+1)), 0666)

	return uint64(i + 1), nil
}
