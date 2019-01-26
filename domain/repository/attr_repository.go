package repository

import (
	"context"

	"github.com/kombu/domain/model"
)

type AttrRepository interface {
	Create(ctx context.Context, inode uint64, mode uint32, name string) error
	FetchById(ctx context.Context, inode int64) (*model.Attr, error)
	Update(ctx context.Context, attr *model.Attr) (*model.Attr, error)
	Delete(ctx context.Context, inode int64) error
	FetchChildrenbyId(ctx context.Context, id int64) (*[]model.Attr, error)
}
