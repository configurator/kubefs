package kube

import (
	"fmt"
	"strings"

	"bazil.org/fuse/fs"
	"github.com/configurator/kubefs/pkg/kfuse"
	"k8s.io/client-go/discovery"
)

func (c *Context) ToDir(kfs *kfuse.KubeFS) *kfuse.Dir {
	return &kfuse.Dir{
		KubeFS:       kfs,
		ReadDirNames: func() ([]string, error) { return c.readDirNames(kfs) },
		LookupNode:   func(name string) (fs.Node, error) { return c.lookupNode(kfs, name) },
	}
}

func SplitGroupVersion(groupVersion string) (string, string) {
	index := strings.LastIndex(groupVersion, "/")
	if index == -1 {
		return "", groupVersion
	} else {
		return groupVersion[0:index], groupVersion[index+1:]
	}
}

func (c *Context) loadResources(kfs *kfuse.KubeFS) error {
	if c.resources != nil {
		// Already loaded
		return nil
	}

	dc, err := discovery.NewDiscoveryClientForConfig(c.restConfig)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resources, err := discovery.ServerPreferredResources(dc)
	if err != nil {
		fmt.Println(err)
		return err
	}

	result := make(map[string]fs.Node)
	for _, list := range resources {
		group, version := SplitGroupVersion(list.GroupVersion)
		for _, r := range list.APIResources {
			// Add missing metadata
			if r.Group == "" {
				r.Group = group
			}
			if r.Version == "" {
				r.Version = version
			}

			if r.Namespaced {
				result[r.Name] = (&NamespacedResource{
					config:     c.config,
					restConfig: c.restConfig,
					context:    c.context,
					Resource:   r,
				}).ToDir(kfs)
			} else {
				result[r.Name] = (&GlobalResource{
					config:     c.config,
					restConfig: c.restConfig,
					context:    c.context,
					Resource:   r,
				}).ToDir(kfs)
			}
		}
	}

	c.resources = result
	return nil
}

func (c *Context) readDirNames(kfs *kfuse.KubeFS) ([]string, error) {
	err := c.loadResources(kfs)
	if err != nil {
		return nil, err
	}

	result := []string{}
	for name := range c.resources {
		result = append(result, name)
	}
	return result, nil
}

func (c *Context) lookupNode(kfs *kfuse.KubeFS, name string) (fs.Node, error) {
	err := c.loadResources(kfs)
	if err != nil {
		return nil, err
	}

	r, _ := c.resources[name]
	// r will be nil if not found, which is what we want
	return r, nil
}
