package repository

import (
	"github.com/kombu/domain/protocol"
)

type DataServer interface {
	Close(protocol.Model) error
	GetWriter(protocol.Model) (ServerWriter, error)
	GetReader(protocol.Model) (ServerReader, error)
}

type serverWriter struct{}

type ServerWriter interface {
	Write(p []byte) (n int, err error)
}

type serverReader struct{}

type ServerReader interface {
	Read(p []byte) (n int, err error)
}
