package kube

import (
	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
)

func (k *Kubernetes) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return k.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return k.lookupNode(kfs, name) },
	}
}

func (k *Kubernetes) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	result := []string{}
	for name := range k.Contexts {
		result = append(result, name)
	}

	return result, nil
}

func (k *Kubernetes) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	context, ok := k.Contexts[name]
	if !ok {
		return nil, nil
	}

	return context.ToDir(kfs), nil
}
