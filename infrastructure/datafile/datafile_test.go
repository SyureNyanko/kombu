package datafile

import (
	"testing"
	"path/filepath"
	 "io/ioutil"
	 "os"
	protocol_model "github.com/kombu/domain/protocol"
)
var (
testfilepath = "testfile.txt"
)

func Prepare() *OpenedDataFileServerImpl{
	apath, _ := filepath.Abs(".")
	s := NewOpenedDataFileServerImpl(apath)

	
	if _, err := os.Stat(testfilepath) ; err != nil {
		content := []byte("hello world\n")
		ioutil.WriteFile(testfilepath, content, os.ModePerm)
	}
	return s
}

/*
type Model struct {
	uri    string
	offset int
	size   int
}

 */
func TestWriter(t *testing.T){
	s := Prepare()
	testData := protocol_model.Model{}
	w, err := s.GetWriter(testData)
	if err != nil {
		t.Error(err)
	}
	w.Write([]byte("Hello World!\n"))
}