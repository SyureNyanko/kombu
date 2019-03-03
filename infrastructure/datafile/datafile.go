package datafile

import (
	"github.com/kombu/domain/data/model"
	protocol_model "github.com/kombu/domain/protocol"
	"os"
	"github.com/rs/xid"
	"path/filepath"
	"sync"
	"errors"
)


/*
 implement Cache file
 */


type writerImpl struct {
	model.Writer
	Fd *os.File
}

type readerImpl struct {
	model.Reader
	Fd *os.File
}

func(w *writerImpl) Write(b []byte){
	w.Fd.Write(b)
}


func(r *readerImpl) Read(b []byte){
	r.Fd.Read(b)
}


type OpenedDataFileServerImpl struct {
	model.DataServer
	localroot string
	p *filemap
}

type filemap struct {
	m *sync.Mutex /* for resource ? */
	mp map[string]*(os.File)
}

func newFilemap() *filemap {
	return &filemap{
		m : new(sync.Mutex),  /* for resource ? */
		mp : map[string](*os.File){}, 
	}
}

func(p *filemap) addDiscriptor(path string, f *os.File) error {
	p.m.Lock()
	defer p.m.Unlock()
	p.mp[path] = f
	return nil
}

func(p *filemap) getDiscriptor(key string) *os.File {
	p.m.Lock()
	defer p.m.Unlock()
	return p.mp[key]
}
func(p *filemap) removeDiscriptor(key string) error {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.mp, key)
	_, exist := p.mp[key]
	if exist {
		return errors.New("Remove Error")
	}
	return nil
}



func NewOpenedDataFileServerImpl(root string) *OpenedDataFileServerImpl {
	fmp := newFilemap()
	return &OpenedDataFileServerImpl{
		model.DataServer{
			Id : "0",
			Backend : model.LocalFileSystem,
		},
		root,
		fmp, 
	}
}

func (c *OpenedDataFileServerImpl) Close(p protocol_model.Model) error {
	f := c.p.getDiscriptor(p.Uri)
	err := f.Close()
	if err != nil {
		return err
	}
	err = c.p.removeDiscriptor(p.Uri)
	return err
}

func (c *OpenedDataFileServerImpl) GetWriter(p protocol_model.Model) (*writerImpl, error) {

	path := getNewFilePath(c.localroot)
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if err := c.p.addDiscriptor(path, f) ; err != nil {
		f.Close()
		return nil, err
	}
	return &writerImpl{Fd: f}, nil
}

func (c *OpenedDataFileServerImpl) GetReader(p protocol_model.Model) (*readerImpl, error) {
	f, err := os.Open(p.Uri)
	if err != nil {
		return nil, err
	}
	if err := c.p.addDiscriptor(p.Uri, f) ; err != nil {
		f.Close()
		return nil, err
	}
	return &readerImpl{Fd: f}, nil
}

func getNewFilePath(root string) string {
	guid := xid.New()
	return filepath.Join(root, guid.String())

}
