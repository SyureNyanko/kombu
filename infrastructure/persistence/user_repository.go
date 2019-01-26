package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kombu/domain/model"
	"github.com/kombu/domain/repository"
)

type AttrConfigure struct {
	uid uint32
	gid uint32
}

// AttrRepositoryImpl Implements repository.AttrRepository
type AttrRepositoryImpl struct {
	dbpath    string
	configure AttrConfigure
}

func NewAttrRepositoryWithSQLite(dbpath string) repository.AttrRepository {
	/*
		ToDO : Connect with upper
	*/
	conf := AttrConfigure{
		100, 100,
	}

	db, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	/*  drop debug only */
	db.DropTableIfExists(&model.Attr{})
	db.CreateTable()
	db.AutoMigrate(&model.Attr{})
	return &AttrRepositoryImpl{dbpath: dbpath, configure: conf}
}

func DBErrorPrint(errors []error) {
	for i, err := range errors {
		fmt.Println(i, err)
	}
}

/*
type Attr struct {
	Id       int
	ParentId int
	Name     string

	Ino       uint64
	Size      uint64
	Blocks    uint64
	Atime     uint64
	Mtime     uint64
	Ctime     uint64
	Atimensec uint32
	Mtimensec uint32
	Ctimensec uint32
	Mode      uint32
	Nlink     uint32
	Uid       uint32
	Gid       uint32
	Rdev      uint32
	Blksize   uint32
	Padding   uint32
}

*/

func (r *AttrRepositoryImpl) CreateAttr(name string, isDir bool) *model.Attr {
	attr := model.Attr{}
	attr.Name = name
	now := time.Now()
	nowunixtime := uint64(now.Unix())
	nowunixnanotime := uint32(now.UnixNano())
	attr.Atime, attr.Mtime, attr.Ctime = nowunixtime, nowunixtime, nowunixtime
	attr.Atimensec, attr.Mtimensec, attr.Ctimensec = nowunixnanotime, nowunixnanotime, nowunixnanotime
	if isDir {
		attr.Mode = fuse.S_IFDIR | 0755
	} else {
		attr.Mode = fuse.S_IFREG | 0644
	}
	attr.Uid = r.configure.uid
	attr.Gid = r.configure.gid
	return &attr
}

func (r *AttrRepositoryImpl) SaveAttr(ctx context.Context, attr *model.Attr) (*model.Attr, error) {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if err := db.Create(&attr).GetErrors(); len(err) != 0 {
		return nil, fmt.Errorf("DB Error")
	}
	return attr, nil
}

func (r *AttrRepositoryImpl) FetchById(ctx context.Context, inode int64) (*model.Attr, error) {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return nil, fmt.Errorf("DB Error")
	}
	defer db.Close()
	var attr model.Attr
	if err := db.First(&attr, "ino = ?", inode).Error; err != nil {
		return nil, err
	}
	return &attr, nil
}

func (r *AttrRepositoryImpl) Update(ctx context.Context, new_attr *model.Attr) (*model.Attr, error) {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var attr model.Attr
	db.First(&attr, "ino = ?", new_attr.Ino)
	db.Model(attr).Updates(new_attr)
	return &attr, nil
}

func (r *AttrRepositoryImpl) Delete(ctx context.Context, inode int64) error {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return err
	}
	defer db.Close()
	var attr model.Attr
	if err := db.First(&attr, "ino = ?", inode).Error; err != nil {
		return fmt.Errorf("DB Error")
	}
	db.Delete(&attr)
	return nil
}

func (r *AttrRepositoryImpl) FetchChildrenbyId(ctx context.Context, parentid int64) (*[]model.Attr, error) {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	attrs := []model.Attr{}
	if err := db.Find(&attrs, "parent_id = ?", parentid).Error; err != nil {
		return nil, fmt.Errorf("DB Error")
	}
	return &attrs, nil

}

func (r *AttrRepositoryImpl) IdisExists(ctx context.Context, id int64) (bool, error) {
	db, err := gorm.Open("sqlite3", r.dbpath)
	if err != nil {
		return false, err
	}
	defer db.Close()
	attrs := []model.Attr{}
	if err := db.Find(&attrs, "id = ?", id).Error; err != nil {
		return false, fmt.Errorf("DB Error")
	}
	if len(attrs) == 0 {
		return false, nil
	}
	return true, nil
}
