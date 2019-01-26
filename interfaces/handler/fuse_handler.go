package handler

import (
	"context"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/usecase"
)

const (
	ROOT_INODE = 1
)

type FuseHandler interface {
	fuse.RawFileSystem
}

type fuseHandler struct {
	r RootHandler
	u usecase.AttrUseCase
}

func NewFuseHandler(mountpoint string, u usecase.AttrUseCase) FuseHandler {
	r := NewRootHandler(mountpoint)
	return &fuseHandler{r, u}
}

func (fs *fuseHandler) Init(*fuse.Server) {

}

func (fs *fuseHandler) String() string {
	return "kombufs"
}

func (fs *fuseHandler) StatFs(header *fuse.InHeader, out *fuse.StatfsOut) fuse.Status {
	if header.NodeId == ROOT_INODE {
		return fs.r.StatFs(header, out)
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) Lookup(header *fuse.InHeader, name string, out *fuse.EntryOut) (code fuse.Status) {
	if header.NodeId == ROOT_INODE {
		return fs.r.Lookup(header, name, out)
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) GetAttr(input *fuse.GetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	if input.NodeId == ROOT_INODE {
		return fs.r.GetAttr(input, out)
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) Access(input *fuse.AccessIn) (code fuse.Status) {
	if input.InHeader.NodeId == ROOT_INODE {
		return fs.r.Access(input)
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) ReadDir(input *fuse.ReadIn, out *fuse.DirEntryList) fuse.Status {
	if input.InHeader.NodeId == ROOT_INODE {
		return fs.r.ReadDir(input, out)
	}
	return fuse.ENOSYS
}

/*
type InHeader struct {
	Length uint32
	Opcode int32
	Unique uint64
	NodeId uint64
	Context
	Padding uint32
}

*/

func (fs *fuseHandler) Create(input *fuse.CreateIn, name string, out *fuse.CreateOut) (code fuse.Status) {
	ctx := context.TODO()
	id := input.InHeader.NodeId
	mode := input.Mode
	if err := fs.u.Create(ctx, id, mode, name); err != nil {
		return fuse.ENOSYS
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) OpenDir(input *fuse.OpenIn, out *fuse.OpenOut) (status fuse.Status) {
	if input.InHeader.NodeId == ROOT_INODE {
		return fs.r.OpenDir(input, out)
	}
	return fuse.ENOSYS
}

func (fs *fuseHandler) Read(input *fuse.ReadIn, buf []byte) (fuse.ReadResult, fuse.Status) {
	return nil, fuse.ENOSYS
}

func (fs *fuseHandler) GetLk(in *fuse.LkIn, out *fuse.LkOut) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) SetLk(in *fuse.LkIn) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) SetLkw(in *fuse.LkIn) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Release(input *fuse.ReleaseIn) {
}

func (fs *fuseHandler) Write(input *fuse.WriteIn, data []byte) (written uint32, code fuse.Status) {
	return 0, fuse.ENOSYS
}

func (fs *fuseHandler) Flush(input *fuse.FlushIn) fuse.Status {
	return fuse.OK
}

func (fs *fuseHandler) Fsync(input *fuse.FsyncIn) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) ReadDirPlus(input *fuse.ReadIn, l *fuse.DirEntryList) fuse.Status {
	return fuse.ENOSYS
}

func (fs *fuseHandler) ReleaseDir(input *fuse.ReleaseIn) {
}

func (fs *fuseHandler) FsyncDir(input *fuse.FsyncIn) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Fallocate(in *fuse.FallocateIn) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Forget(nodeID, nlookup uint64) {
}

func (fs *fuseHandler) GetXAttrSize(header *fuse.InHeader, attr string) (size int, code fuse.Status) {
	return 0, fuse.ENOSYS
}

func (fs *fuseHandler) GetXAttrData(header *fuse.InHeader, attr string) (data []byte, code fuse.Status) {
	return nil, fuse.ENOATTR
}

func (fs *fuseHandler) SetXAttr(input *fuse.SetXAttrIn, attr string, data []byte) fuse.Status {
	return fuse.ENOSYS
}

func (fs *fuseHandler) ListXAttr(header *fuse.InHeader) (data []byte, code fuse.Status) {
	return nil, fuse.ENOSYS
}

func (fs *fuseHandler) RemoveXAttr(header *fuse.InHeader, attr string) fuse.Status {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Readlink(header *fuse.InHeader) (out []byte, code fuse.Status) {
	return nil, fuse.ENOSYS
}

func (fs *fuseHandler) Mknod(input *fuse.MknodIn, name string, out *fuse.EntryOut) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Mkdir(input *fuse.MkdirIn, name string, out *fuse.EntryOut) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Unlink(header *fuse.InHeader, name string) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Rmdir(header *fuse.InHeader, name string) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Symlink(header *fuse.InHeader, pointedTo string, linkName string, out *fuse.EntryOut) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Rename(input *fuse.RenameIn, oldName string, newName string) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) Link(input *fuse.LinkIn, name string, out *fuse.EntryOut) (code fuse.Status) {
	return fuse.ENOSYS
}
func (fs *fuseHandler) Open(input *fuse.OpenIn, out *fuse.OpenOut) (status fuse.Status) {
	return fuse.OK
}

func (fs *fuseHandler) SetAttr(input *fuse.SetAttrIn, out *fuse.AttrOut) (code fuse.Status) {
	return fuse.ENOSYS
}

func (fs *fuseHandler) SetDebug(dbg bool) {
}
