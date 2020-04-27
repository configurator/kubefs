package kubefs

import (
	"context"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type KubeFS struct {
	Uid uint32
	Gid uint32
}

func (f *KubeFS) Root() (fs.Node, error) {
	root := &RootDir{KubeFS: f}
	return root, nil
}

func (mo *KubeFS) defaultAttr(directory bool, ctx context.Context, attr *fuse.Attr) error {
	if directory {
		attr.Mode = os.ModeDir | 0444
	} else {
		attr.Mode = 0444
	}
	attr.Uid = mo.Uid
	attr.Gid = mo.Gid
	return nil
}
