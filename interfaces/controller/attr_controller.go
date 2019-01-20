package controller

import (
    "github.com/kombu/usecase"
)

const (
	ATTR_VALID = 0
	ATTR_NSEC = 0
)


type AttrController struct {
	*root
}

func (c *AttrController) GetAttr(input *fuse.GetAttrIn, out *fuse.AttrOut) error {
	nodeid := input.InHeader.NodeId
	ctx := context.TODO()
	attr, err := GetAttr(ctx, nodeid)
	if err != nil {
		return err
	}
	out = &AttrOut{
		Attrvallid : ATTR_VALID,
		AttrValidNsec : ATTR_NSEC,
		Attr : fuse.Attr{
			Ino       : attr.Ino,
			Size      : attr.Size,
			Blocks    : attr.Blocks,
			Atime     : attr.Atime,
			Mtime     : attr.Mtime,
			Ctime     : attr.Ctime,
			Atimensec : attr.Atimensec,
			Mtimensec : attr.Mtimensec,
			Ctimensec : attr.Ctimensec,
			Mode      : attr.Mode,
			Nlink     : attr.Nlink,
			Owner : fuse.Owner{
				Uid: attr.Uid,
				Gif : attr.Gid,
			}
			Rdev    : attr.Rdev,
			Blksize : attr.Blksize,
			Padding : attr.Padding,
		}
	}
}