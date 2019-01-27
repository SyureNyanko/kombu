package service

import (
	"github.com/kombu/domain/model"
	"github.com/hanwen/go-fuse/fuse"
)

type AttrService interface {
	TransformFuseToModel(inFuseAttr *fuse.Attr, outModelAttr *model.Attr) error
	TransformModelToFuse(outModelAttr *model.Attr, inFuseAttr *fuse.Attr) error
}

type attrService struct {
	is model.InodeServer
}

type NewAttrService(is model.InodeServer) AttrService {
	return &attrService{is}
}


func (s *attrService) TransformFuseToModel(inFuseAttr *fuse.Attr, outModelAttr *model.Attr) error {

}

func  (s *attrService) TransformModelToFuse(outModelAttr *model.Attr, inFuseAttr *fuse.Attr) error {
}