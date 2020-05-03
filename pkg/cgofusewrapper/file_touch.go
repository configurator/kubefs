package cgofusewrapper

import (
	"log"

	"github.com/billziss-gh/cgofuse/fuse"
)

// Utimens changes the access and modification times of a file.
func (fs *FS) Utimens(path string, tmsp []fuse.Timespec) int {
	log.Printf("fs.Utimens(%v, %#v)\n", path, tmsp)

	// `touch` changes the file time, so we support that operation so touch doesn't fail when
	// creating new files. We don't actually need to do anything - just return a success code after
	// which `touch` will write the empty file in the nonexistant case.
	return 0
}
