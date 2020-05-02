package cgofusewrapper

type File interface {
	ReadEntireContents() ([]byte, error)
	Write([]byte) error
	Delete() error
}

type BaseFile struct{}

var _ Node = (*BaseFile)(nil)

func (f *BaseFile) Attr(stat *Stat) (FileType, FilePermissions, error) {
	return FileType_File, FilePermissions_Read, nil
}
