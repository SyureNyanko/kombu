package handler

import (
	"github.com/hanwen/go-fuse/fuse"
	"log"
)

type RootHandler interface {
	GetAttr(*fuse.GetAttrIn, *fuse.AttrOut) fuse.Status
	StatFs( *fuse.InHeader, *fuse.StatfsOut) fuse.Status 
	Lookup( *fuse.InHeader,  string, *fuse.EntryOut) fuse.Status
	OpenDir( *fuse.OpenIn, *fuse.OpenOut) ( fuse.Status) 
	Access(*fuse.AccessIn)fuse.Status
	ReadDir(*fuse.ReadIn,*fuse.DirEntryList) fuse.Status 
}

type rootHandler struct {

}

func NewRootHandler(mountpoint string) RootHandler {
	return &rootHandler{

	}
}

func (fs *rootHandler) GetAttr(input *fuse.GetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	*out = fuse.AttrOut {
		Attr: fuse.Attr {
			Ino : ROOT_INODE,
			Mode: fuse.S_IFDIR | 0755,
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
	*out = fuse.OpenOut {
	}
	return fuse.OK
}

func (fs *rootHandler) Access(input *fuse.AccessIn) (code fuse.Status) {
	return fuse.OK
}

func (fs *rootHandler) ReadDir(input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	d := []byte{}
	*out = *fuse.NewDirEntryList(d, 0)
	/* https://github.com/hanwen/go-fuse/blob/291273cb8ce0f139636a6fd7414be3c7e2de6288/fuse/direntry.go */
	return fuse.OK
}
