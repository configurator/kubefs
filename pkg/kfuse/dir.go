package kfuse

import (
	"context"
	"fmt"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dir struct {
	*KubeFS

	ReadDirNames func() ([]string, error)
	LookupNode   func(name string) (fs.Node, error)
}

var _ fs.Node = (*Dir)(nil)
var _ fs.NodeRequestLookuper = (*Dir)(nil)
var _ fs.HandleReadDirAller = (*Dir)(nil)

func (d *Dir) Attr(ctx context.Context, attr *fuse.Attr) error {
	return d.defaultAttr(os.ModeDir, attr)
}

func (d *Dir) Lookup(ctx context.Context, request *fuse.LookupRequest, response *fuse.LookupResponse) (fs.Node, error) {
	if d.LookupNode == nil {
		return nil, fmt.Errorf("LookupNode not implemented")
	}

	result, err := d.LookupNode(request.Name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fuse.ENOENT
	}

	return result, nil
}

func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	result := []fuse.Dirent{}
	if d.ReadDirNames == nil {
		return result, nil
	}

	names, err := d.ReadDirNames()
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		result = append(result, fuse.Dirent{
			Name: name,
		})
	}

	return result, nil
}
