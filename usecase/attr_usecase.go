package usecase

import (
	"context"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/domain/model"
	"github.com/kombu/domain/repository"
	"github.com/kombu/interfaces/controller"
)

type AttrUseCase interface {
	Create(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.EntryOut, error)
	GetAttr(ctx context.Context, id int64) (*model.Attr, error)
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

func (interactor *attrInteractor) Create(ctx context.Context, header *fuse.InHeader, mode uint32, name string) (*fuse.EntryOut, error) {
	inode, err := interactor.InodeServer.IssueId()
	if err != nil {
		return nil, err
	}
	modelAttr, err := interactor.AttrRepository.Create(ctx, header.NodeId /* parent nodeid */, inode, mode, name)
	if err != nil {
		return nil, err
	}
	fuseAttr := interactor.Controller.ModelToFuse(modelAttr)
	entryout := interactor.Controller.FuseAttrToEntryOut(fuseAttr)
	return entryout, err
}

func (interactor *attrInteractor) GetAttr(ctx context.Context, id int64) (*model.Attr, error) {
	a, err := interactor.AttrRepository.FetchById(ctx, id)
	return a, err
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
