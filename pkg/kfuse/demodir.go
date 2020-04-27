package kfuse

import (
	"bazil.org/fuse/fs"
)

func DemoDir(kfs *KubeFS) fs.Node {
	return &Dir{
		KubeFS: kfs,

		ReadDirNames: func() ([]string, error) {
			return []string{
				"thefile",
				"link",
			}, nil
		},

		LookupNode: func(name string) (fs.Node, error) {
			if name == "thefile" {
				return &File{
					KubeFS: kfs,
					ReadContents: func() ([]byte, error) {
						return []byte("Example file contents here\n"), nil
					},
				}, nil
			}
			if name == "link" {
				return &Symlink{
					KubeFS: kfs,
					Target: "thefile",
				}, nil
			}

			return nil, nil
		},
	}
}
