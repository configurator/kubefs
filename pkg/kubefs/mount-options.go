package kubefs

import (
	"context"
	"os"

	"bazil.org/fuse"
)

type MountOptions struct {
	Uid uint32
	Gid uint32
}

func (mo *MountOptions) defaultAttr(ctx context.Context, attr *fuse.Attr) error {
	attr.Mode = os.ModeDir | 0444
	attr.Uid = mo.Uid
	attr.Gid = mo.Gid
	return nil
}
