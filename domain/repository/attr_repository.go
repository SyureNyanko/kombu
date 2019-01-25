package repository

import (
	"context"

	"github.com/kombu/domain/model"
)

type AttrRepository interface {
	CreateFileAttr(name string) *model.Attr
	CreateDirectoryAttr(name string) *model.Attr

	SaveAttr(ctx context.Context, attr *model.Attr) (*model.Attr, error)
	FetchById(ctx context.Context, inode int64) (*model.Attr, error)
	Update(ctx context.Context, attr *model.Attr) (*model.Attr, error)
	Delete(ctx context.Context, inode int64) error
	FetchChildrenbyId(ctx context.Context, id int64) (*[]model.Attr, error)
}
