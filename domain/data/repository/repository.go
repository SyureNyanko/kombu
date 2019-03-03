package repository

import (
	"github.com/kombu/domain/protocol"
	"github.com/kombu/domain/data/repository"
)

type DataServer interface {
	Close(protocol.Model) error
	GetWriter(protocol.Model) (*model.Writer, error)
	GetReader(protocol.Model) (*model.Reader, error)
}



type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}