package cgofusewrapper

import (
	"log"
	"sync"

	"github.com/billziss-gh/cgofuse/fuse"
	"github.com/configurator/kubefs/pkg/cgofusewrapper/errno"
)

const FileHandleValue = ^uint64(0)

type FileHandle struct {
	File     File
	Data     []byte
	Mutex    sync.Mutex
	released bool
}

func (fs *FS) getFileHandle(path string) (*FileHandle, int) {
	handle := fs.Handles.Get(path)
	if handle == nil {
		log.Printf("Handle %s not found\n", path)
		// No such handle id
		return nil, errno.EIO
	}

	if result, ok := handle.(*FileHandle); ok {
		return result, 0
	}

	// A handle exists with that id, but it's not a FileHandle
	log.Printf("Handle %s not a FileHandle: %#v\n", path, handle)
	return nil, errno.EIO
}

func (fs *FS) openOrCreate(path string, fi *fuse.FileInfo_t, canCreate bool) int {
	node, err := fs.findNode(path)
	if err != nil {
		return handleError(err)
	}

	if file, ok := node.(File); ok {
		err := fs.Handles.CreateOrIncrement(path, func() (interface{}, error) {
			data, err := file.ReadEntireContents()
			if err != nil && !canCreate {
				return nil, err
			}
			// errors can happen in read if the file doesn't exist for example
			// in those cases we ignore the error start with an empty file
			return &FileHandle{
				File: file,
				Data: data,
			}, nil
		})
		if err != nil {
			return handleError(err)
		}

		fi.DirectIo = true
		fi.NonSeekable = true
		fi.KeepCache = false
		fi.Fh = FileHandleValue
		return 0
	}

	if _, ok := node.(Dir); ok {
		return errno.EISDIR
	}
	return errno.EOPNOTSUPP
}

func (fs *FS) OpenEx(path string, fi *fuse.FileInfo_t) int {
	return fs.openOrCreate(path, fi, true)
}

func (fs *FS) CreateEx(path string, _ uint32, fi *fuse.FileInfo_t) int {
	return fs.openOrCreate(path, fi, false)
}

func (fs *FS) Truncate(path string, size int64, fh uint64) (errn int) {
	handle, e := fs.getFileHandle(path)
	if e != 0 {
		return e
	}

	handle.Mutex.Lock()
	defer handle.Mutex.Unlock()
	if handle.released {
		// Write after close
		return errno.EIO
	}

	oldData := handle.Data
	handle.Data = make([]byte, size)
	copy(handle.Data, oldData)
	return 0
}

func (fs *FS) Write(path string, buff []byte, ofst int64, fh uint64) int {
	handle, errn := fs.getFileHandle(path)
	if errn != 0 {
		return errno.EIO
	}

	handle.Mutex.Lock()
	defer handle.Mutex.Unlock()
	if handle.released {
		// Write after close
		return errno.EIO
	}

	writeFileHandle(handle, buff, ofst)
	return len(buff)
}

func writeFileHandle(handle *FileHandle, buff []byte, offset int64) {
	// Writing is a bit complicated, because we may need to extend the buffer or not, so we split
	// it into two operations
	l := int64(len(handle.Data))
	if offset == l {
		// Quick path for appending to the file
		handle.Data = append(handle.Data, buff...)
		return
	}

	// Extend the output with zeroes if needed, then copy the buffer into it
	writeEnd := offset + int64(len(buff))
	if writeEnd > l {
		handle.Data = append(handle.Data, make([]byte, writeEnd-l)...)
	}

	copy(handle.Data[offset:], buff)
}

func (fs *FS) Release(path string, fh uint64) int {
	handle, errn := fs.getFileHandle(path)
	if errn != 0 {
		return errn
	}

	// The lock here makes sure we're not writing _while_ the buffer is being modified
	// some write calls may be delayed after the Release ends due to locks, and these
	// will error out
	handle.Mutex.Lock()
	defer handle.Mutex.Unlock()

	err := handle.File.Write(handle.Data)
	if err != nil {
		return handleError(err)
	}
	return 0
}
