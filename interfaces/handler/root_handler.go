package handler

import (
	"github.com/hanwen/go-fuse/fuse"
	"log"
)

type RootHandler interface {
	GetAttr(*fuse.GetAttrIn, *fuse.AttrOut) fuse.Status
	StatFs( *fuse.InHeader, *fuse.StatfsOut) fuse.Status 
	Lookup( *fuse.InHeader,  string, *fuse.EntryOut) fuse.Status
}

type rootHandler struct {

}

func NewRootHandler(mountpoint string) RootHandler {
	return &rootHandler{

	}
}

func (fs *rootHandler) GetAttr(input *fuse.GetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	out = &fuse.AttrOut {
		Attr: fuse.Attr {
			Ino : ROOT_INODE,
		},
	}
	
	log.Println("RootGetAttr")
	return fuse.OK
}

func (fs *rootHandler) StatFs(header *fuse.InHeader, out *fuse.StatfsOut) fuse.Status {
	log.Println("StatFs")
	return fuse.OK
}

func (fs *rootHandler) Lookup(header *fuse.InHeader, name string, out *fuse.EntryOut) (code fuse.Status) {
	log.Println("Lookup")
	return fuse.OK
}