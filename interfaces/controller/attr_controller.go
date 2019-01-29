package controller

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/domain/model"
)

type AttrController interface {
	ModelToFuse(m *model.Attr) *fuse.Attr
	FuseAttrToEntryOut(fa *fuse.Attr) *fuse.EntryOut
	FuseAttrToCreateOut(fa *fuse.Attr) *fuse.CreateOut
}

type attrController struct{}

func NewAttrController() AttrController {
	return &attrController{}
}

/*
type Attr struct {
	Ino         uint64
	Size        uint64
	Blocks      uint64
	Atime       uint64
	Mtime       uint64
	Ctime       uint64
	Crtime_     uint64 // OS X
	Atimensec   uint32
	Mtimensec   uint32
	Ctimensec   uint32
	Crtimensec_ uint32 // OS X
	Mode        uint32
	Nlink       uint32
	Owner
	Rdev   uint32
	Flags_ uint32 //  OS X
}
*/
/*
type Attr struct {
	Id       uint64
	ParentId uint64
	Name     string

	Ino       uint64
	Size      uint64
	Blocks    uint64
	Atime     uint64
	Mtime     uint64
	Ctime     uint64
	Atimensec uint32
	Mtimensec uint32
	Ctimensec uint32
	Mode      uint32
	Nlink     uint32
	Uid       uint32
	Gid       uint32
	Rdev      uint32
	Blksize   uint32
	Padding   uint32
}

*/
func (a *attrController) ModelToFuse(m *model.Attr) *fuse.Attr {
	return &fuse.Attr{
		Ino:       m.Ino,
		Size:      m.Size,
		Blocks:    m.Blocks,
		Atime:     m.Atime,
		Mtime:     m.Mtime,
		Ctime:     m.Ctime,
		Atimensec: m.Atimensec,
		Mode:      m.Mode,
		Nlink:     m.Nlink,
		Owner: fuse.Owner{
			Uid: m.Uid,
			Gid: m.Gid,
		},
		Rdev: m.Rdev,
	}
}

/*
type EntryOut struct {
	NodeId         uint64
	Generation     uint64
	EntryValid     uint64
	AttrValid      uint64
	EntryValidNsec uint32
	AttrValidNsec  uint32
	Attr
}


*/

func (a *attrController) FuseAttrToEntryOut(fa *fuse.Attr) *fuse.EntryOut {
	/*
		TODO : Generation ...
	*/
	return &fuse.EntryOut{
		NodeId: fa.Ino,
		Attr:   *fa,
	}
}

/*
type CreateOut struct {
	EntryOut
	OpenOut
}
*/

func (a *attrController) FuseAttrToCreateOut(fa *fuse.Attr) *fuse.CreateOut {
	/*
		TODO : Generation ...
	*/
	entry := a.FuseAttrToEntryOut(fa)
	openout := fuse.OpenOut{
		Fh:        0,
		OpenFlags: fuse.FOPEN_KEEP_CACHE,
		Padding:   0,
	}
	return &fuse.CreateOut{
		EntryOut: *entry,
		OpenOut:  openout,
	}
}
