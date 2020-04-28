package kube

import (
	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
)

func (n *NamespacedResource) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return n.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return n.lookupNode(kfs, name) },
	}
}

func (n *NamespacedResource) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	return []string{}, nil
}

func (n *NamespacedResource) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	return nil, nil
}
