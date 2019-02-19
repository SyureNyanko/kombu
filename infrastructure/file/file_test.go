package file

import (
	"testing"

	"github.com/hanwen/go-fuse/fuse"
    "github.com/kombu/domain/file/model"
)

func TestAddAndRetrieveDescriptor(t *testing.T) {
	ts := NewOpenedFileServer()
	/* OpenDir */
	de := ts.NewDirEntry()
	entrylist := fuse.NewDirEntryList([]byte{}, 0)
	entrylist.AddDirEntry(fuse.DirEntry{
		Mode: fuse.S_IFDIR,
		Name: "test",
		Ino:  1,
	})
	de.AddOneEntry(entrylist)

	entrylist2 := fuse.NewDirEntryList([]byte{}, 0)
	entrylist2.AddDirEntry(fuse.DirEntry{
		Mode: fuse.S_IFDIR,
		Name: "test",
		Ino:  1,
	})
	de.AddOneEntry(entrylist2)
	d := model.DirEntry{
		Entries: []*fuse.DirEntryList{entrylist, entrylist2},
		Count:   2,
	}
	desc, err := ts.Register(&d)
	if err != nil {
		t.Error("Register Error")
	}
	direntry, err := ts.Retrieve(desc)
	if err != nil {
		t.Error("Register Error")
	}
	dl := direntry.RetrieveOneEntry()
	t.Logf("d : %+v", dl)

	dl2 := direntry.RetrieveOneEntry()
	t.Logf("d2 : %+v", dl2)

	dl3 := direntry.RetrieveOneEntry()
	if dl3 != nil {
		t.Error("not nil Error")
	}
}
