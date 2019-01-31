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
	GetChildren(ctx context.Context, id int64) (*[]model.Attr, error)
	DeleteAttr(ctx context.Context, id int64) error
	UpdateAttr(ctx context.Context, id int64, a *model.Attr) error
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

func (interactor *attrInteractor) GetChildren(ctx context.Context, id int64) (*[]model.Attr, error) {
	attrs, err := interactor.AttrRepository.FetchChildrenbyId(ctx, id)
	return attrs, err
}

func (interactor *attrInteractor) DeleteAttr(ctx context.Context, id int64) error {
	err := interactor.AttrRepository.Delete(ctx, id)
	return err
}

func (interactor *attrInteractor) UpdateAttr(ctx context.Context, id int64, a *model.Attr) error {
	_, err := interactor.AttrRepository.Update(ctx, a)
	return err
}
