package usecase

import (
	"context"

	"github.com/kombu/domain/model"
	"github.com/kombu/domain/repository"
)

type AttrUseCase interface {
	CreateNewAttr(ctx context.Context, a *model.Attr) error
	GetAttr(ctx context.Context, id int64) (*model.Attr, error)
	GetChildren(ctx context.Context, id int64) (*[]model.Attr, error)
	DeleteAttr(ctx context.Context, id int64) error
	UpdateAttr(ctx context.Context, id int64, a *model.Attr) error
}

type attrInteractor struct {
	AttrRepository repository.AttrRepository
}

func NewAttrInteractor(r repository.AttrRepository) AttrUseCase {
	return &attrInteractor{r}
}

func (interactor *attrInteractor) CreateNewAttr(ctx context.Context, a *model.Attr) error {
	_, err := interactor.AttrRepository.Create(ctx, a)
	return err
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
