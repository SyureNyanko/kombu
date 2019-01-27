package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/kombu/infrastructure/inode"
	"github.com/kombu/infrastructure/persistence"
	"github.com/kombu/interfaces/handler"
	"github.com/kombu/usecase"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Printf("usage: %s MOUNTPOINT\n", path.Base(os.Args[0]))
		fmt.Printf("\noptions:\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	mountPoint := flag.Arg(0)

	sqlRepository := persistence.NewAttrRepositoryWithSQLite("test.sqlite")
	inodeServerRepository := inode.NewInodeServerImpl("inodefile")
	usecase := usecase.NewAttrInteractor(sqlRepository, inodeServerRepository)
	kombufs := handler.NewFuseHandler(mountPoint, usecase)
	fmt.Println(mountPoint)
	mountOpts := fuse.MountOptions{}
	mountOpts.Debug = true
	state, _ := fuse.NewServer(kombufs, mountPoint, &mountOpts)
	state.Serve()
}
