package kubefs

import (
	"context"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type RootDir struct {
	*KubeFS
}

var _ fs.Node = (*RootDir)(nil)
var _ fs.NodeRequestLookuper = (*RootDir)(nil)
var _ fs.HandleReadDirAller = (*RootDir)(nil)

func (d *RootDir) Attr(ctx context.Context, attr *fuse.Attr) error {
	return d.defaultAttr(true, ctx, attr)
}

func (d *RootDir) Lookup(ctx context.Context, request *fuse.LookupRequest, response *fuse.LookupResponse) (fs.Node, error) {
	path := request.Name
	if path == "thefile" {
		return &TheFile{KubeFS: d.KubeFS}, nil
	}
	return nil, fuse.ENOENT
}

func (d *RootDir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return []fuse.Dirent{
		fuse.Dirent{
			Inode: 15,
			Name:  "thefile",
			Type:  fuse.DT_File,
		},
	}, nil
}
