package util

import (
	"syscall"
	"github.com/hanwen/go-fuse/fuse"
)
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
func GetRootAttr(root string) (*fuse.Attr, error ){
	var st syscall.Stat_t
	err := syscall.Stat(root, &st)
	if err != nil {
		return nil, err
	}
	var a1 fuse.Attr
	a1.FromStat(&st)
	a1.Ino = 2
	return &a1, nil
}