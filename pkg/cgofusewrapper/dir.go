package cgofusewrapper

import (
	"log"

	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

type Dir interface {
	Node
	List() ([]string, error)
	Get(string) (Node, error)
}

type BaseDir struct{}

var _ Node = (*BaseDir)(nil)

func (d *BaseDir) Attr(stat *Stat) (FileType, FilePermissions, error) {
	return FileType_Directory, FilePermissions_ReadExecute, nil
}

// Readdir reads a directory.
func (fs *FS) Readdir(path string,
	fill func(name string, stat *fuse.Stat_t, ofst int64) bool,
	ofst int64,
	fh uint64) int {

	log.Printf("fs.Readdir(%v, callback, %#x, fh)\n", path, ofst)

	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	dir, ok := node.(Dir)
	if !ok {
		return errno.ENOTDIR
	}

	children, err := dir.List()
	if err != nil {
		return handleError(err)
	}

	fill(".", nil, 0)
	fill("..", nil, 0)
	for _, name := range children {
		fill(name, nil, 0)
	}
	return 0
}
