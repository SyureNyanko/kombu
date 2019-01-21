package handler

import (
	"github.com/hanwen/go-fuse/fuse"
	"fmt"
)

type RootHandler interface {
	GetAttr(*fuse.GetAttrIn, *fuse.AttrOut) (fuse.Status)
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
	
	fmt.Println("RootGetAttr")
	return fuse.OK
}