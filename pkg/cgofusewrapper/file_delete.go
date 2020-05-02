package cgofusewrapper

import (
	"log"

	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

// Unlink removes a file.
func (fs *FS) Unlink(path string) int {
	log.Printf("fs.Unlink(%v)\n", path)

	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	if file, ok := node.(File); ok {
		err := file.Delete()
		if err != nil {
			return handleError(err)
		}
		return 0
	}

	return errno.EOPNOTSUPP
}
