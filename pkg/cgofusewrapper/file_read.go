package cgofusewrapper

import (
	"log"

	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

// Read reads data from a file.
func (fs *FS) Read(path string, buff []byte, offset int64, fh uint64) int {
	log.Printf("fs.Read(%v, buffer, %#x, fh)\n", path, offset)

	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	file, ok := node.(File)
	if !ok {
		if _, ok := node.(Dir); ok {
			return errno.EISDIR
		}
		return errno.EIO
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
