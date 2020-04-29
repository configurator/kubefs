package cgofusewrapper

import "github.com/billziss-gh/cgofuse/fuse"

type Stat fuse.Stat_t

type FileType uint32
type FilePermissions uint32

const (
	FileType_Directory FileType = fuse.S_IFDIR
	FileType_Link      FileType = fuse.S_IFLNK
	FileType_File      FileType = fuse.S_IFREG

	FilePermissions_Read             FilePermissions = 0444
	FilePermissions_ReadExecute      FilePermissions = 0555
	FilePermissions_ReadWrite        FilePermissions = 0664
	FilePermissions_ReadWriteExecute FilePermissions = 0775
)

type Node interface {
	Attr(stat *Stat) (FileType, FilePermissions, error)
}
