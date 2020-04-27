package kubefs

import (
	"context"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type TheFile struct {
	*KubeFS
}

var _ fs.Node = (*TheFile)(nil)
var _ fs.NodeOpener = (*TheFile)(nil)

func (f *TheFile) Attr(ctx context.Context, attr *fuse.Attr) error {
	return f.KubeFS.defaultAttr(false, ctx, attr)
}

func (f *TheFile) Open(ctx context.Context, req *fuse.OpenRequest, resp *fuse.OpenResponse) (fs.Handle, error) {
	resp.Flags = fuse.OpenDirectIO | fuse.OpenNonSeekable
	return &TheFileHandle{KubeFS: f.KubeFS}, nil
}

type TheFileHandle struct {
	*KubeFS
}

var _ fs.Handle = (*TheFileHandle)(nil)
var _ fs.HandleReader = (*TheFileHandle)(nil)

var contents = []byte("Some example file contents\n")

func (fh *TheFileHandle) Read(ctx context.Context, req *fuse.ReadRequest, resp *fuse.ReadResponse) error {
	sliceStart := req.Offset
	sliceEnd := req.Offset + int64(req.Size)
	if sliceEnd > int64(len(contents)) {
		sliceEnd = int64(len(contents))
	}
	resp.Data = contents[sliceStart:sliceEnd]
	return nil
}
