package kfuse

import (
	"os"
	"os/user"
	"strconv"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type KubeFS struct {
	Uid uint32
	Gid uint32

	RootDir fs.Node
}

func (f *KubeFS) ReadCurrentUser() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	uid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return err
	}
	gid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return err
	}

	f.Uid = uint32(uid)
	f.Gid = uint32(gid)
	return nil
}

func (f *KubeFS) Root() (fs.Node, error) {
	return f.RootDir, nil
}

func (mo *KubeFS) defaultAttr(modeFlags os.FileMode, attr *fuse.Attr) error {
	attr.Mode = modeFlags | 0444
	attr.Uid = mo.Uid
	attr.Gid = mo.Gid
	return nil
}
