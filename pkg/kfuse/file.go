package kubefs

import (
	"context"
	"fmt"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type File struct {
	*KubeFS

	ReadContents func() ([]byte, error)
}

var _ fs.Node = (*File)(nil)
var _ fs.NodeOpener = (*File)(nil)

func (f *File) Attr(ctx context.Context, attr *fuse.Attr) error {
	return f.KubeFS.defaultAttr(0, attr)
}

func (f *File) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	if f.ReadContents == nil {
		return nil, fmt.Errorf("Read is not implemented for this file")
	}

	contents, err := f.ReadContents()
	if err != nil {
		return nil, err
	}

	resp.Flags = fuse.OpenDirectIO | fuse.OpenNonSeekable
	return &FileHandle{File: f, Contents: contents}, nil
}

type FileHandle struct {
	*File

	Contents []byte
}

var _ fs.Handle = (*FileHandle)(nil)
var _ fs.HandleReader = (*FileHandle)(nil)

func (fh *FileHandle) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	sliceStart := req.Offset
	sliceEnd := req.Offset + int64(req.Size)
	if sliceEnd > int64(len(fh.Contents)) {
		sliceEnd = int64(len(fh.Contents))
	}
	resp.Data = fh.Contents[sliceStart:sliceEnd]
	return nil
}
