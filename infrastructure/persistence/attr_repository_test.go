package persistence

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kombu/domain/model"
)

const (
	testdb = "testdata/test.sqlite3"
)

func cleanUpData() {
	os.Remove(testdb)
}

func createTestAttr() *model.Attr {
	m := model.Attr{Name: "Mikan"}
	return &m
}

func dumpDB() *[]model.Attr {
	db, err := gorm.Open("sqlite3", testdb)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	attrs := []model.Attr{}
	db.Find(&attrs)
	for a := range attrs {
		fmt.Printf("Attr : %+v", a)
	}
	return &attrs
}

func TestCRUD(t *testing.T) {
	cleanUpData()
	a := createTestAttr()
	r := NewAttrRepositoryWithSQLite(testdb)
	ctx := context.TODO()
	a.Ino = 2
	a.Name = "Yuzu"
	a.Id = 2
	a.ParentId = 10
	_, err := r.Create(ctx, a)
	a_ret, err := r.FetchById(ctx, 2)
	if err != nil {
		t.Fatal("Fetch Error")
	}
	if a_ret.Name != "Yuzu" || a_ret.Ino != 2 {
		t.Errorf("Fetch Equality Error %+v", a_ret)
	}
	a_ret.Name = "Lemon"
	_, err = r.Update(ctx, a_ret)
	if err != nil {
		t.Errorf("Update Error %+v", err)
	}
	a_ret, err = r.FetchById(ctx, 2)
	if a_ret.Name != "Lemon" || err != nil {
		t.Errorf("Update Equality Error %+v", err)
	}

	/*  add another attr */
	a_2 := createTestAttr()
	a_2.Name = "Grape"
	a_2.Id = 3
	a_2.ParentId = 10
	_, err = r.Create(ctx, a_2)
	attrs, err := r.FetchChildrenbyId(ctx, 10)
	if err != nil {
		t.Fatal("FetchChildrenbyId Error")
	}
	if len(*attrs) != 2 {
		t.Error("FetchChildrenbyId Sum Error", len(*attrs))
	}
	for _, attr := range *attrs {
		t.Log(attr)
		if !((attr.Name == "Lemon" && attr.Id == 2) || (attr.Name == "Grape" && attr.Id == 3)) {
			t.Error("FetchChildrenbyId Contents Error")
		}
	}
	err = r.Delete(ctx, 2)
	if err != nil {
		t.Error("Delete Error")
	}
	a_ret, err = r.FetchById(ctx, 2)
	if a_ret != nil {
		t.Errorf("Delete Failed Error %+v, %s", a_ret)
	}

}
