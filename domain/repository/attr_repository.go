package repository

import (
	"context"

	"github.com/kombu/domain/model"
)

type AttrRepository interface {
	Create(ctx context.Context, parentid uint64, inode uint64, mode uint32, name string) (*model.Attr, error)
	FetchById(ctx context.Context, inode uint64) (*model.Attr, error)
	Update(ctx context.Context, attr *model.Attr) (*model.Attr, error)
	Delete(ctx context.Context, inode int64) error
	FetchChildrenbyId(ctx context.Context, id uint64) (*[]model.Attr, error)
}

type InodeServer interface {
	IssueId() (uint64, error)
}

