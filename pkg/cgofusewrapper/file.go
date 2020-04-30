package cgofusewrapper

import (
	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

type File interface {
	ReadEntireContents() ([]byte, error)
	Delete() error
}

type BaseFile struct{}

var _ Node = (*BaseFile)(nil)

func (f *BaseFile) Attr(stat *Stat) (FileType, FilePermissions, error) {
	return FileType_File, FilePermissions_Read, nil
}

func (fs *FS) Read(path string, buff []byte, offset int64, fh uint64) int {
	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	file, ok := node.(File)
	if !ok {
		if _, ok := node.(Dir); ok {
			return errno.EISDIR
		}
		return errno.EUNKNOWN
	}

	contents, err := file.ReadEntireContents()
	if err != nil {
		return handleError(err)
	}

	end := offset + int64(len(buff))
	if end > int64(len(contents)) {
		end = int64(len(contents))
	}
	if end < offset {
		return 0
	}

	return copy(buff, contents[offset:end])
}
