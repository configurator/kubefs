package kfuse

import (
	"context"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Symlink struct {
	*KubeFS
	Target string
}

var _ fs.Node = (*Symlink)(nil)
var _ fs.NodeReadlinker = (*Symlink)(nil)

func (f *Symlink) Attr(ctx context.Context, attr *fuse.Attr) error {
	return f.KubeFS.defaultAttr(os.ModeSymlink, attr)
}

func (f *Symlink) Readlink(ctx context.Context, request *fuse.ReadlinkRequest) (string, error) {
	return f.Target, nil
}
