package kube

import (
	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
)

func (c *Context) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return c.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return c.lookupNode(kfs, name) },
	}
}

func (c *Context) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	result := []string{"TODO"}

	return result, nil
}

func (c *Context) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	return nil, nil
}
