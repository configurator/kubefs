package cgofusewrapper

import (
	"log"

	"github.com/billziss-gh/cgofuse/fuse"

	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

// Init is called when the file system is created.
func (fs *FS) Init() {
	log.Printf("(unimplemented) fs.Init()\n")
}

// Destroy is called when the file system is destroyed.
func (fs *FS) Destroy() {
	log.Printf("(unimplemented) fs.Destroy()\n")
}

// Statfs gets file system statistics.
func (fs *FS) Statfs(path string, stat *fuse.Statfs_t) int {
	log.Printf("(unimplemented) fs.Statfs(%v, stat)\n", path)
	return errno.ENOSYS
}

// Mknod creates a file node.
func (fs *FS) Mknod(path string, mode uint32, dev uint64) int {
	log.Printf("(unimplemented) fs.Mknod(%v)\n", path)
	return errno.ENOSYS
}

// Mkdir creates a directory.
func (fs *FS) Mkdir(path string, mode uint32) int {
	log.Printf("(unimplemented) fs.Mkdir(%v)\n", path)
	return errno.ENOSYS
}

// Rmdir removes a directory.
func (fs *FS) Rmdir(path string) int {
	log.Printf("(unimplemented) fs.Rmdir(%v)\n", path)
	return errno.ENOSYS
}

// Link creates a hard link to a file.
func (fs *FS) Link(oldpath string, newpath string) int {
	log.Printf("(unimplemented) fs.Link(%v, %v)\n", oldpath, newpath)
	return errno.ENOSYS
}

// Symlink creates a symbolic link.
func (fs *FS) Symlink(target string, newpath string) int {
	log.Printf("(unimplemented) fs.Symlink(%v, %v)\n", target, newpath)
	return errno.ENOSYS
}

// Readlink reads the target of a symbolic link.
func (fs *FS) Readlink(path string) (int, string) {
	log.Printf("(unimplemented) fs.Readlink(%v)\n", path)
	return errno.ENOSYS, ""
}

// Rename renames a file.
func (fs *FS) Rename(oldpath string, newpath string) int {
	log.Printf("(unimplemented) fs.Rename(%v, %v)\n", oldpath, newpath)
	return errno.ENOSYS
}

// Chmod changes the permission bits of a file.
func (fs *FS) Chmod(path string, mode uint32) int {
	log.Printf("(unimplemented) fs.Chmod(%v, %#o)\n", path, mode)
	return errno.ENOSYS
}

// Chown changes the owner and group of a file.
func (fs *FS) Chown(path string, uid uint32, gid uint32) int {
	log.Printf("(unimplemented) fs.Chmod(%v, %d, %d)\n", path, uid, gid)
	return errno.ENOSYS
}

// Access checks file access permissions.
func (fs *FS) Access(path string, mask uint32) int {
	log.Printf("(unimplemented) fs.Access(%v, %#o)\n", path, mask)
	return errno.ENOSYS
}

// Create creates and opens a file.
// The flags are a combination of the fuse.O_* constants.
func (fs *FS) Create(path string, flags int, mode uint32) (int, uint64) {
	log.Printf("(unimplemented) fs.Create(%v, %#x, %#o)\n", path, flags, mode)
	return errno.ENOSYS, ^uint64(0)
}

// Open opens a file.
// The flags are a combination of the fuse.O_* constants.
func (fs *FS) Open(path string, flags int) (int, uint64) {
	log.Printf("(unimplemented) fs.Open(%v, %#x)\n", path, flags)
	return errno.ENOSYS, ^uint64(0)
}

// Flush flushes cached file data.
func (fs *FS) Flush(path string, fh uint64) int {
	log.Printf("(unimplemented) fs.Flush(%v, fh)\n", path)
	return errno.ENOSYS
}

// Fsync synchronizes file contents.
func (fs *FS) Fsync(path string, datasync bool, fh uint64) int {
	log.Printf("(unimplemented) fs.Fsync(%v, %v, fh)\n", path, datasync)
	return errno.ENOSYS
}

/*
// Lock performs a file locking operation.
func (fs *FS) Lock(path string, cmd int, lock *Lock_t, fh uint64) int {
	return errno.ENOSYS
}
*/

// Opendir opens a directory.
func (fs *FS) Opendir(path string) (int, uint64) {
	log.Printf("(unimplemented) fs.Opendir(%v)\n", path)
	return errno.ENOSYS, ^uint64(0)
}

// Releasedir closes an open directory.
func (fs *FS) Releasedir(path string, fh uint64) int {
	log.Printf("(unimplemented) fs.Releasedir(%v, fh)\n", path)
	return errno.ENOSYS
}

// Fsyncdir synchronizes directory contents.
func (fs *FS) Fsyncdir(path string, datasync bool, fh uint64) int {
	log.Printf("(unimplemented) fs.Fsyncdir(%v, %v, fh)\n", path, datasync)
	return errno.ENOSYS
}

// Setxattr sets extended attributes.
func (fs *FS) Setxattr(path string, name string, value []byte, flags int) int {
	log.Printf("(unimplemented) fs.Setxattr(%v, %v, %v, %#x)\n", path, name, value, flags)
	return errno.ENOSYS
}

// Getxattr gets extended attributes.
func (fs *FS) Getxattr(path string, name string) (int, []byte) {
	log.Printf("(unimplemented) fs.Getxattr(%v, %v)\n", path, name)
	return errno.ENOSYS, nil
}

// Removexattr removes extended attributes.
func (fs *FS) Removexattr(path string, name string) int {
	log.Printf("(unimplemented) fs.Removexattr(%v, %v)\n", path, name)
	return errno.ENOSYS
}

// Listxattr lists extended attributes.
func (fs *FS) Listxattr(path string, fill func(name string) bool) int {
	log.Printf("(unimplemented) fs.Listxattr(%v, callback)\n", path)
	return errno.ENOSYS
}
