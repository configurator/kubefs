package cgofusewrapper

import (
	"log"
	"strings"

	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

type FS struct {
	Handles Handles
	Root    Dir
}

var _ fuse.FileSystemInterface = (*FS)(nil)
var _ fuse.FileSystemOpenEx = (*FS)(nil)

func (fs *FS) findNode(path string) (Node, error) {
	node := fs.Root.(Node)

	parts := strings.Split(path, "/")
	for i, part := range parts {
		if part != "" {
			dir, ok := node.(Dir)
			if !ok {
				partialPath := strings.Join(parts[0:i+1], "/")
				return nil, &ErrorNotADirectory{Path: partialPath}
			}

			child, err := dir.Get(part)
			if err != nil {
				return nil, err
			}
			node = child
		}
	}

	return node, nil
}

func handleError(err error) int {
	if err, ok := err.(FuseError); ok {
		log.Println(err.Error())
		return err.ErrorCode()
	}

	log.Printf("Unknown error: %s\n", err)
	return errno.EUNKNOWN
}

func fullStat(node Node, stat *fuse.Stat_t) error {
	t, p, err := node.Attr((*Stat)(stat))
	stat.Mode = uint32(t) | uint32(p)
	return err
}

// Getattr gets file attributes.
func (fs *FS) Getattr(path string, stat *fuse.Stat_t, fh uint64) int {
	log.Printf("fs.Getattr(%v, stat, fh)\n", path)

	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	err = fullStat(node, stat)
	if err != nil {
		return handleError(err)
	}
	return 0
}
