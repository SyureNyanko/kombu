package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/domain/model"
	"github.com/kombu/domain/repository"
	"github.com/kombu/interfaces/controller"
)

type AttrUseCase interface {
	Create(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.CreateOut, error)
	Mknod(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.EntryOut, error)
	GetAttr(ctx context.Context, header *fuse.InHeader) (*fuse.AttrOut, error)
	DeleteAttr(ctx context.Context, id int64) error
	UpdateAttr(ctx context.Context, id int64, a *model.Attr) error
	OpenDir(ctx context.Context, header *fuse.InHeader) (*fuse.OpenOut, error)
	ReadDir(ctx context.Context, header *fuse.InHeader, size uint32, offset uint64) (*fuse.DirEntryList, error)
}

type attrInteractor struct {
	AttrRepository repository.AttrRepository
	InodeServer    repository.InodeServer
	Controller     controller.AttrController
}

func NewAttrInteractor(r repository.AttrRepository, i repository.InodeServer) AttrUseCase {
	c := controller.NewAttrController()
	return &attrInteractor{r, i, c}
}

func (interactor *attrInteractor) Create(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.CreateOut, error) {
	inode, err := interactor.InodeServer.IssueId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	modelAttr, err := interactor.AttrRepository.Create(ctx, header.NodeId /* parent nodeid */, inode, mode, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fuseAttr := interactor.Controller.ModelToFuse(modelAttr)
	fmt.Printf("create modelAttr %+v", *fuseAttr)
	createout := interactor.Controller.FuseAttrToCreateOut(fuseAttr)
	fmt.Printf("create entryout %+v", *createout)
	return createout, err
}

func (interactor *attrInteractor) Mknod(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.EntryOut, error) {
	inode, err := interactor.InodeServer.IssueId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	modelAttr, err := interactor.AttrRepository.Create(ctx, header.NodeId /* parent nodeid */, inode, mode, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fuseAttr := interactor.Controller.ModelToFuse(modelAttr)
	fmt.Printf("modelAttr %+v", *fuseAttr)
	entryout := interactor.Controller.FuseAttrToEntryOut(fuseAttr)
	fmt.Printf("entryout %+v", *entryout)
	return entryout, err
}

func (interactor *attrInteractor) GetAttr(ctx context.Context, header *fuse.InHeader) (*fuse.AttrOut, error) {
	modelAttr, err := interactor.AttrRepository.FetchById(ctx, header.NodeId)
	fuseAttr := interactor.Controller.ModelToFuse(modelAttr)
	return &fuse.AttrOut{Attr: *fuseAttr}, err
}

func (interactor *attrInteractor) DeleteAttr(ctx context.Context, id int64) error {
	err := interactor.AttrRepository.Delete(ctx, id)
	return err
}

func (interactor *attrInteractor) UpdateAttr(ctx context.Context, id int64, a *model.Attr) error {
	_, err := interactor.AttrRepository.Update(ctx, a)
	return err
}

/* TODO: implement Open/OpenDir(issue file descripter?) */
func (interactor *attrInteractor) OpenDir(ctx context.Context, header *fuse.InHeader) (*fuse.OpenOut, error) {
	return &fuse.OpenOut{
		Fh: 3,
	}, nil
}

func (interactor *attrInteractor) ReadDir(ctx context.Context, header *fuse.InHeader, size uint32, offset uint64) (*fuse.DirEntryList, error) {
	buf := make([]byte, size)
	entrylist := fuse.NewDirEntryList(buf, offset)
	attrs, err := interactor.AttrRepository.FetchChildrenbyId(ctx, header.NodeId)
	log.Println("------")
	for _, a := range *attrs {
		log.Printf("Name : %s", a.Name)
		entrylist.AddDirEntry(fuse.DirEntry{
			Mode: a.Mode,
			Name: a.Name,
			Ino:  a.Ino,
		})
	}
	log.Println("------")
	return entrylist, err
}
