package handler

import (
	"log"

	"github.com/hanwen/go-fuse/fuse"
)

type RootHandler interface {
	GetAttr(*fuse.GetAttrIn, *fuse.AttrOut) fuse.Status
	StatFs(*fuse.InHeader, *fuse.StatfsOut) fuse.Status
	Lookup(*fuse.InHeader, string, *fuse.EntryOut) fuse.Status
	OpenDir(*fuse.OpenIn, *fuse.OpenOut) fuse.Status
	Access(*fuse.AccessIn) fuse.Status
}

type rootHandler struct {
}

func NewRootHandler(mountpoint string) RootHandler {
	return &rootHandler{}
}

func (fs *rootHandler) GetAttr(input *fuse.GetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	*out = fuse.AttrOut{
		Attr: fuse.Attr{
			Ino:  ROOT_INODE,
			Mode: fuse.S_IFDIR | 0777,
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

func (fs *rootHandler) OpenDir(input *fuse.OpenIn, out *fuse.OpenOut) (status fuse.Status) {
	*out = fuse.OpenOut{}
	return fuse.OK
}

func (fs *rootHandler) Access(input *fuse.AccessIn) (code fuse.Status) {
	return fuse.OK
}
