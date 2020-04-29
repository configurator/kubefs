package main

import (
	"github.com/billziss-gh/cgofuse/fuse"
	f "github.com/configurator/kubefs/pkg/cgofusewrapper"
)

func main() {
	h := fuse.NewFileSystemHost(&f.FS{
		Root: &MyDir{},
	})
	h.Mount("", []string{"/home/configurator/testmount"})
}

type MyDir struct {
	f.BaseDir
}

var _ f.Dir = (*MyDir)(nil)

func NewFile() *MyFile {
	return &MyFile{}
}

func (*MyDir) List() (map[string]f.Node, error) {
	return map[string]f.Node{
		"hello": NewFile(),
		"file":  NewFile(),
	}, nil
}

func (*MyDir) Get(name string) (f.Node, error) {
	if name == "hello" || name == "file" {
		return NewFile(), nil
	} else {
		return nil, nil
	}
}

type MyFile struct {
	f.BaseFile
}

func (*MyFile) ReadEntireContents() ([]byte, error) {
	return []byte("Hello world!\n"), nil
}
